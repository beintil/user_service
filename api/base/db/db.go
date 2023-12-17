package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/beintil/user_service/api/base/config"
	"github.com/beintil/user_service/pkg/logger"
	lue "github.com/beintil/user_service/pkg/lu_error"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TxFunc(ctx context.Context, opts *sql.TxOptions, fn func(tx *sqlx.Tx) lue.IError) func(*sqlx.DB) lue.IError {
	return func(db *sqlx.DB) lue.IError {
		tx, err := db.BeginTxx(ctx, opts)
		if err != nil {
			return lue.New(err.Error(), lue.InternalServerError)
		}
		e := fn(tx)
		if e != nil {
			tx.Rollback()
			return e
		}
		tx.Commit()
		return nil
	}
}

func GetDB() *sqlx.DB {
	cfg := config.GetConfig().DB
	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.Port))
	if err != nil {
		logger.Error(lue.New("ERROR CONNECTIONS DATABASE "+err.Error(), lue.ErrorConnectDb))
	}
	err = db.Ping()
	if err != nil {
		logger.Error(lue.New("ERROR PING "+err.Error(), lue.ErrorConnectDb))
	}
	return db
}
