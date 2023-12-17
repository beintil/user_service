package core

import (
	"context"
	"database/sql"
	"github.com/beintil/user_service/api/base/db"
	"github.com/beintil/user_service/api/client"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/dimonrus/gohelp"
	"github.com/dimonrus/gosql"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	*client.User
}

func (m *User) SetPrimary(id int64) *User {
	m.Id = gohelp.Ptr(id)
	return m
}

func (m *User) RemovePassword() *User {
	m.Password = nil
	return m
}

func NewUser() *User {
	return &User{client.NewUser()}
}

func (m *User) Create(sqlDb *sqlx.DB) lue.IError {
	var options = &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return db.TxFunc(ctx, options, func(tx *sqlx.Tx) lue.IError {
		return m.CreateTX(tx)
	})(sqlDb)
}

func (m *User) CreateTX(tx *sqlx.Tx) lue.IError {
	if m == nil {
		return lue.New("user is empty", lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	if m.Password == nil {
		return lue.New("password is empty", lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	validate := validator.New()
	err := validate.Struct(*m)
	if err != nil {
		return lue.New(err.Error(), lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(*m.Password), bcrypt.DefaultCost)
	if err != nil {
		return lue.New(err.Error(), lue.InternalServerError).SetHttpCode(http.StatusInternalServerError)
	}

	query, values, _ := gosql.PGSQL(m.insertUserSql(password))
	err = tx.QueryRow(query, values...).Scan(&m.Person.UserId)
	if err != nil {
		return lue.New(err.Error(), lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}

	query, values, _ = gosql.PGSQL(m.insertPersonSql())
	_, err = tx.Exec(query, values...)
	if err != nil {
		return lue.New(err.Error(), lue.ErrorDatabaseQuery).SetHttpCode(http.StatusInternalServerError)
	}
	return nil
}

func (m *User) insertUserSql(password []byte) *gosql.Insert {
	s := gosql.NewInsert().Into(m.Table())
	s.Columns().Add(m.Columns()[1:]...)
	m.SetCreatedAt()
	s.Columns().Arg(password, m.Email, m.CreatedAt, m.UpdatedAt)
	s.Returning().Add("id")
	return s
}

func (m *User) insertPersonSql() *gosql.Insert {
	s := gosql.NewInsert().Into(m.Person.Table())
	m.Person.SetCreatedAt()
	s.Columns().Add(m.Person.Columns()[1:]...)
	s.Columns().Arg(m.Person.Values()[1:]...)
	return s
}

func (m *User) Load(sqlDb *sqlx.DB) lue.IError {
	var options = &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return db.TxFunc(ctx, options, func(tx *sqlx.Tx) lue.IError {
		return m.loadTx(tx)
	})(sqlDb)
}

func (m *User) loadTx(sql *sqlx.Tx) lue.IError {
	query, values, _ := gosql.PGSQL(m.loadUserSql())
	scanValues := append([]any{}, m.Values()...)
	scanValues = append(scanValues, m.Person.Values()...)
	err := sql.QueryRow(query, values...).Scan(scanValues...)
	if err != nil {
		return lue.New(err.Error(), lue.BadRequest).SetHttpCode(http.StatusBadRequest)
	}
	return nil
}

func (m *User) GetPassword(sql *sqlx.Tx) ([]byte, lue.IError) {
	var password []byte
	row := sql.QueryRow(`SELECT password FROM "user" WHERE id = $1`, &m.Id)
	if row.Err() != nil {
		return nil, lue.New(row.Err().Error(), lue.InternalServerError).SetHttpCode(http.StatusInternalServerError)
	}
	err := row.Scan(&password)
	if err != nil {
		return nil, lue.New(err.Error(), lue.ErrorDatabaseArgument).SetHttpCode(http.StatusInternalServerError)
	}
	return password, nil
}

func (m *User) loadUserSql() gosql.ISQL {
	s := gosql.NewSelect().From(`"user" u`)
	s.Columns().Add(SetAlias("u", m.Columns()...)...)
	s.Where().AddExpression(`u.id = ? OR u.email = ?`, m.Id, m.Email)
	s.Relate(`LEFT JOIN person ON person.user_id = u.id`)
	s.Columns().Add(SetAlias("person", m.Person.Columns()...)...)
	return s
}

func UsersSearch(sql *sqlx.DB, form client.UserSearchForm) (items []*client.User, e lue.IError) {
	query, values, _ := gosql.PGSQL(getSearchSql(form))
	rows, err := sql.Query(query, values...)
	if err != nil {
		e = lue.New(err.Error(), lue.ErrorDatabaseQuery).SetHttpCode(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		items = append(items, client.NewUser())
		err = rows.Scan(items[len(items)-1].ValuesWithPerson()...)
		if err != nil {
			e = lue.New(err.Error(), lue.ErrorDatabaseArgument).SetHttpCode(http.StatusInternalServerError)
			return
		}
	}
	return
}

func getSearchSql(form client.UserSearchForm) *gosql.Select {
	var user = NewUser()
	s := gosql.NewSelect().From(`"user" u`)
	s.Columns().Add("u.id", "u.email", "u.created_at", "u.updated_at")
	cond := form.PrepareCondition()
	if !cond.IsEmpty() {
		s.Where().Replace(cond)
	}
	s.Relate("JOIN person p ON u.id = p.user_id")
	s.Columns().Add(SetAlias("p", user.Person.Columns()...)...)
	return s
}

func (m *User) Verify(sql *sqlx.DB) bool {
	user := NewUser()
	user.Email = m.Email
	e := user.Load(sql)
	if e != nil {
		return false
	}
	correct := m.checkPasswordMatchAndHashing([]byte(*user.Password))
	if !correct {
		return false
	}
	m.User = user.User
	return true
}

func (m *User) checkPasswordMatchAndHashing(hashPassword []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(*m.Password))
	return err == nil
}
