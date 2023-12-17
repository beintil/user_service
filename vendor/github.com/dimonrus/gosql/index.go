package gosql

import "strings"

// CREATE [ UNIQUE ] INDEX [ CONCURRENTLY ] [ [ IF NOT EXISTS ] name ] ON [ ONLY ] table_name [ USING method ]
//
//	( { column_name | ( expression ) } [ COLLATE collation ] [ opclass [ ( opclass_parameter = value [, ... ] ) ] ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ] [, ...] )
//	[ INCLUDE ( column_name [, ...] ) ]
//	[ NULLS [ NOT ] DISTINCT ]
//	[ WITH ( storage_parameter [= value] [, ... ] ) ]
//	[ TABLESPACE tablespace_name ]
//	[ WHERE predicate ]
type Index struct {
	// UNIQUE
	unique bool
	// CONCURRENTLY
	concurrently bool
	// IF NOT EXISTS
	ifNotExists bool
	// name of index
	name string
	// ONLY
	only bool
	// table_name
	tableName string
	// USING method
	using string
	// index expression
	expression expression
	// INCLUDE ( column_name [, ...] )
	include expression
	// WITH ( storage_parameter [= value] [, ... ] )
	with expression
	// TABLESPACE tablespace_name
	tablespace string
	// WHERE predicate
	where Condition
	// generate name automatically
	autoName bool
	// is nulls not distinct
	nullsNotDistinct bool
}

// Using method
func (i *Index) Using(using string) *Index {
	i.using = using
	return i
}

// Include add include columns
func (i *Index) Include(column ...string) *Index {
	i.include.Add(column...)
	return i
}

// With add with params
func (i *Index) With(param ...string) *Index {
	i.with.Add(param...)
	return i
}

// Name set index name
func (i *Index) Name(name string) *Index {
	i.name = name
	return i
}

// Table set table name
func (i *Index) Table(name string) *Index {
	i.tableName = name
	return i
}

// AutoName generate name on render
func (i *Index) AutoName() *Index {
	i.autoName = true
	return i
}

// Concurrently create index
func (i *Index) Concurrently() *Index {
	i.concurrently = true
	return i
}

// AutoName generate name on render
func (i *Index) getAutoName() string {
	b := strings.Builder{}
	if i.tableName != "" {
		b.WriteString(i.tableName + "_")
	}
	var word strings.Builder
	for _, c := range i.expression.String(EnumDelimiter) {
		if 'a' <= c && c <= 'z' {
			word.WriteRune(c)
		} else if 'A' <= c && c <= 'Z' {
			word.WriteRune(c)
		} else {
			b.WriteString(word.String() + "_")
			word.Reset()
		}
	}
	if word.Len() > 0 {
		b.WriteString(word.String() + "_")
	}
	if i.unique {
		b.WriteString("uidx")
	} else {
		b.WriteString("idx")
	}
	return b.String()
}

// Expression index expression
func (i *Index) Expression() *expression {
	return &i.expression
}

// TableSpace set table space
func (i *Index) TableSpace(space string) *Index {
	i.tablespace = space
	return i
}

// Where get where condition
func (i *Index) Where() *Condition {
	return &i.where
}

// NullsNotDistinct set null not distinct
func (i *Index) NullsNotDistinct() *Index {
	i.nullsNotDistinct = true
	return i
}

// Unique set unique
func (i *Index) Unique() *Index {
	i.unique = true
	return i
}

// IfNotExists set if not exists
func (i *Index) IfNotExists() *Index {
	i.ifNotExists = true
	return i
}

// IsEmpty check if empty
func (i *Index) IsEmpty() bool {
	return i == nil || (i.name == "" &&
		i.tableName == "" &&
		i.tablespace == "" &&
		i.using == "" &&
		i.expression.Len() > 0 &&
		i.include.Len() == 0 &&
		i.with.Len() == 0 &&
		i.where.IsEmpty())
}

// String render index query
func (i *Index) String() string {
	if i.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("CREATE")
	if i.unique {
		if i.nullsNotDistinct {
			b.WriteString(" UNIQUE NULLS NOT DISTINCT")
		} else {
			b.WriteString(" UNIQUE")
		}
	}
	b.WriteString(" INDEX")
	if i.concurrently {
		b.WriteString(" CONCURRENTLY")
	}
	if i.ifNotExists {
		b.WriteString(" IF NOT EXISTS")
	}
	if i.name != "" {
		b.WriteString(" " + i.name)
	} else if i.autoName {
		b.WriteString(" " + i.getAutoName())
	}
	b.WriteString(" ON")
	if i.only {
		b.WriteString(" ONLY")
	}
	if i.tableName != "" {
		b.WriteString(" " + i.tableName)
	}
	if i.using != "" {
		b.WriteString(" USING " + i.using)
	}
	if i.expression.Len() > 0 {
		b.WriteString(" (" + i.expression.String(", ") + ")")
	}
	if i.include.Len() > 0 {
		b.WriteString(" INCLUDE (" + i.include.String(", ") + ")")
	}
	if i.with.Len() > 0 {
		b.WriteString(" WITH (" + i.with.String(", ") + ")")
	}
	if i.tablespace != "" {
		b.WriteString(" TABLESPACE " + i.tablespace)
	}
	if !i.where.IsEmpty() {
		b.WriteString(" WHERE " + i.where.String())
	}
	return b.String() + ";"
}

// SQL common sql interface
func (i *Index) SQL() (query string, params []any, returning []any) {
	query = i.String()
	return
}

// CreateIndex new index
func CreateIndex(arg ...string) *Index {
	var i Index
	if len(arg) > 0 {
		i.tableName = arg[0]
		if len(arg) > 1 {
			i.expression.Add(arg[1:]...)
		}
	}
	return &i
}
