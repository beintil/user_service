package gosql

import "strings"

// ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//
//	action [, ... ]
//
// ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//
//	RENAME [ COLUMN ] column_name TO new_column_name
//
// ALTER TABLE [ IF EXISTS ] [ ONLY ] name [ * ]
//
//	RENAME CONSTRAINT constraint_name TO new_constraint_name
//
// ALTER TABLE [ IF EXISTS ] name
//
//	RENAME TO new_name
//
// ALTER TABLE [ IF EXISTS ] name
//
//	SET SCHEMA new_schema
//
// ALTER TABLE ALL IN TABLESPACE name [ OWNED BY role_name [, ... ] ]
//
//	SET TABLESPACE new_tablespace [ NOWAIT ]
//
// ALTER TABLE [ IF EXISTS ] name
//
//	ATTACH PARTITION partition_name { FOR VALUES partition_bound_spec | DEFAULT }
//
// ALTER TABLE [ IF EXISTS ] name
//
//	DETACH PARTITION partition_name [ CONCURRENTLY | FINALIZE ]
//
// and partition_bound_spec is:
//
// IN ( partition_bound_expr [, ...] ) |
// FROM ( { partition_bound_expr | MINVALUE | MAXVALUE } [, ...] )
//
//	TO ( { partition_bound_expr | MINVALUE | MAXVALUE } [, ...] ) |
//
// WITH ( MODULUS numeric_literal, REMAINDER numeric_literal )
//
// and column_constraint is:
//
// [ CONSTRAINT constraint_name ]
// { NOT NULL |
//
//	NULL |
//	CHECK ( expression ) [ NO INHERIT ] |
//	DEFAULT default_expr |
//	GENERATED ALWAYS AS ( generation_expr ) STORED |
//	GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence_options ) ] |
//	UNIQUE [ NULLS [ NOT ] DISTINCT ] index_parameters |
//	PRIMARY KEY index_parameters |
//	REFERENCES reftable [ ( refcolumn ) ] [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ]
//	  [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
//
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
//
// and table_constraint is:
//
// [ CONSTRAINT constraint_name ]
// { CHECK ( expression ) [ NO INHERIT ] |
//
//	UNIQUE [ NULLS [ NOT ] DISTINCT ] ( column_name [, ... ] ) index_parameters |
//	PRIMARY KEY ( column_name [, ... ] ) index_parameters |
//	EXCLUDE [ USING index_method ] ( exclude_element WITH operator [, ... ] ) index_parameters [ WHERE ( predicate ) ] |
//	FOREIGN KEY ( column_name [, ... ] ) REFERENCES reftable [ ( refcolumn [, ... ] ) ]
//	  [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ] [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
//
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
//
// and table_constraint_using_index is:
//
//	[ CONSTRAINT constraint_name ]
//	{ UNIQUE | PRIMARY KEY } USING INDEX index_name
//	[ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
//
// index_parameters in UNIQUE, PRIMARY KEY, and EXCLUDE constraints are:
//
// [ INCLUDE ( column_name [, ... ] ) ]
// [ WITH ( storage_parameter [= value] [, ... ] ) ]
// [ USING INDEX TABLESPACE tablespace_name ]
//
// exclude_element in an EXCLUDE constraint is:
//
// { column_name | ( expression ) } [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ]
//
// referential_action in a FOREIGN KEY/REFERENCES constraint is:
//
// { NO ACTION | RESTRICT | CASCADE | SET NULL [ ( column_name [, ... ] ) ] | SET DEFAULT [ ( column_name [, ... ] ) ] }
type Alter struct {
	// ordered expression
	ordered orderedExpression
	// partition bound
	bound partitionBound
	// action
	actions []*alterTableAction
}

// Action get action
func (a *Alter) Action() *alterTableAction {
	action := &alterTableAction{}
	a.actions = append(a.actions, action)
	return action
}

// DetachPartition name
func (a *Alter) DetachPartition(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("DETACH PARTITION ", name))
	return a
}

// DetachPartitionConcurrently name
func (a *Alter) DetachPartitionConcurrently(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("DETACH PARTITION ", name, " CONCURRENTLY"))
	return a
}

// DetachPartitionFinalize name
func (a *Alter) DetachPartitionFinalize(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("DETACH PARTITION ", name, " FINALIZE"))
	return a
}

// AttachDefaultPartition name
func (a *Alter) AttachDefaultPartition(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("ATTACH PARTITION ", name, " DEFAULT"))
	return a
}

// AttachPartition name
func (a *Alter) AttachPartition(name string) *partitionBound {
	a.ordered.Add(3, a.ordered.Concat("ATTACH PARTITION ", name))
	return &a.bound
}

// AllInTableSpace name
func (a *Alter) AllInTableSpace(name string) *Alter {
	a.ordered.Add(0, a.ordered.Concat("ALL IN TABLESPACE ", name))
	return a
}

// OwnedBy role
func (a *Alter) OwnedBy(role ...string) *Alter {
	a.ordered.Add(1, a.ordered.Concat("OWNED BY ", strings.Join(role, ", ")))
	return a
}

// SetTableSpace name
func (a *Alter) SetTableSpace(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("SET TABLESPACE ", name))
	return a
}

// SetTableSpaceNoWait name
func (a *Alter) SetTableSpaceNoWait(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("SET TABLESPACE ", name, " NOWAIT"))
	return a
}

// SetSchema table
func (a *Alter) SetSchema(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("SET SCHEMA ", name))
	return a
}

// Rename table
func (a *Alter) Rename(name string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("RENAME TO ", name))
	return a
}

// RenameConstraint rename constraint
func (a *Alter) RenameConstraint(old, new string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("RENAME CONSTRAINT ", old, " TO ", new))
	return a
}

// RenameColumn rename column
func (a *Alter) RenameColumn(old, new string) *Alter {
	a.ordered.Add(3, a.ordered.Concat("RENAME COLUMN ", old, " TO ", new))
	return a
}

// IfExists set if exists
func (a *Alter) IfExists() *Alter {
	a.ordered.Add(0, a.ordered.Concat("IF EXISTS"))
	return a
}

// Only set
func (a *Alter) Only() *Alter {
	a.ordered.Add(1, a.ordered.Concat("ONLY"))
	return a
}

// Name set name
func (a *Alter) Name(name string) *Alter {
	a.ordered.Add(2, a.ordered.Concat(name))
	return a
}

// IsEmpty check if empty
func (a *Alter) IsEmpty() bool {
	return a == nil || a.ordered.IsEmpty() && len(a.actions) == 0
}

// String render alter table query
func (a *Alter) String() string {
	if a.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("ALTER TABLE ")
	b.WriteString(a.ordered.String())
	if !a.bound.IsEmpty() {
		b.WriteString(" FOR VALUES " + a.bound.String())
	}
	for i := range a.actions {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(" " + a.actions[i].String())
	}
	return b.String() + ";"
}

// SQL common sql interface
func (a *Alter) SQL() (query string, params []any, returning []any) {
	query = a.String()
	params = a.ordered.GetArguments()
	for i := range a.actions {
		params = append(params, a.actions[i].GetArguments()...)
	}
	return
}

// AlterTable table constructor
func AlterTable(args ...string) *Alter {
	alter := &Alter{}
	if len(args) > 0 {
		alter.ordered.Add(2, alter.ordered.Concat(args[0]))
	}
	return alter
}
