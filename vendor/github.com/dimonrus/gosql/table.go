package gosql

import "strings"

const (
	// ReferencesMatchFull MATCH FULL
	ReferencesMatchFull = "FULL"
	// ReferencesMatchPartial MATCH PARTIAL
	ReferencesMatchPartial = "PARTIAL"
	// ReferencesMatchSimple MATCH SIMPLE
	ReferencesMatchSimple = "SIMPLE"

	// CheckNoInherit NO INHERIT
	CheckNoInherit = "NO INHERIT"

	// GeneratedStored STORED
	GeneratedStored = "STORED"
	// GeneratedAlways ALWAYS
	GeneratedAlways = "ALWAYS"
	// GeneratedByDefault BY DEFAULT
	GeneratedByDefault = "BY DEFAULT"

	// Deferrable DEFERRABLE
	Deferrable = "DEFERRABLE"
	// NotDeferrable NOT DEFERRABLE
	NotDeferrable = "NOT DEFERRABLE"

	// InitiallyDeferred DEFERRED
	InitiallyDeferred = "DEFERRED"
	// InitiallyImmediate IMMEDIATE
	InitiallyImmediate = "IMMEDIATE"

	// ActionNoAction NO ACTION
	ActionNoAction = "NO ACTION"
	// ActionCascade CASCADE
	ActionCascade = "CASCADE"
	// ActionRestrict RESTRICT
	ActionRestrict = "RESTRICT"
	// ActionSetNull SET NULL
	ActionSetNull = "SET NULL"
	// ActionSetDefault SET DEFAULT
	ActionSetDefault = "SET DEFAULT"

	// LikeIncluding INCLUDING
	LikeIncluding = "INCLUDING"
	// LikeExcluding EXCLUDING
	LikeExcluding = "EXCLUDING"
	// LikeComments COMMENTS
	LikeComments = "COMMENTS"
	// LikeCompression COMPRESSION
	LikeCompression = "COMPRESSION"
	// LikeConstraints CONSTRAINTS
	LikeConstraints = "CONSTRAINTS"
	// LikeDefaults DEFAULTS
	LikeDefaults = "DEFAULTS"
	// LikeGenerated GENERATED
	LikeGenerated = "GENERATED"
	// LikeIdentity IDENTITY
	LikeIdentity = "IDENTITY"
	// LikeIndexes INDEXES
	LikeIndexes = "INDEXES"
	// LikeStatistics STATISTICS
	LikeStatistics = "STATISTICS"
	// LikeStorage STORAGE
	LikeStorage = "STORAGE"
	// LikeAll ALL
	LikeAll = "ALL"

	// WithWithoutOIDS WITHOUT OIDS
	WithWithoutOIDS = "WITHOUT OIDS"

	// PartitionByRange RANGE
	PartitionByRange = "RANGE"
	// PartitionByList LIST
	PartitionByList = "LIST"
	// PartitionByHash HASH
	PartitionByHash = "HASH"

	// PartitionBoundFromMin MINVALUE
	PartitionBoundFromMin = "MINVALUE"
	// PartitionBoundFromMax MAXVALUE
	PartitionBoundFromMax = "MAXVALUE"
	// PartitionBoundWithModulus MODULUS
	PartitionBoundWithModulus = "MODULUS"
	// PartitionBoundWithRemainder REMAINDER
	PartitionBoundWithRemainder = "REMAINDER"

	// PartitionOfWithOptions WITH OPTIONS
	PartitionOfWithOptions = "WITH OPTIONS"
)

// Table create table query builder
// CREATE [ [ GLOBAL | LOCAL ] { TEMPORARY | TEMP } | UNLOGGED ] TABLE [ IF NOT EXISTS ] table_name ( [
//
//	{ column_name data_type [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ]
//	  | table_constraint
//	  | LIKE source_table [ like_option ... ] }
//	  [, ... ]
//
// ] )
// [ INHERITS ( parent_table [, ... ] ) ]
// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
// [ USING method ]
// [ WITH ( storage_parameter [= value] [, ... ] ) | WITHOUT OIDS ]
// [ ON COMMIT { PRESERVE ROWS | DELETE ROWS | DROP } ]
// [ TABLESPACE tablespace_name ]
//
// CREATE [ [ GLOBAL | LOCAL ] { TEMPORARY | TEMP } | UNLOGGED ] TABLE [ IF NOT EXISTS ] table_name
//
//	  OF type_name [ (
//	{ column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//	  | table_constraint }
//	  [, ... ]
//
// ) ]
// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
// [ USING method ]
// [ WITH ( storage_parameter [= value] [, ... ] ) | WITHOUT OIDS ]
// [ ON COMMIT { PRESERVE ROWS | DELETE ROWS | DROP } ]
// [ TABLESPACE tablespace_name ]
//
// CREATE [ [ GLOBAL | LOCAL ] { TEMPORARY | TEMP } | UNLOGGED ] TABLE [ IF NOT EXISTS ] table_name
//
//	  PARTITION OF parent_table [ (
//	{ column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//	  | table_constraint }
//	  [, ... ]
//
// ) ] { FOR VALUES partition_bound_spec | DEFAULT }
// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
// [ USING method ]
// [ WITH ( storage_parameter [= value] [, ... ] ) | WITHOUT OIDS ]
// [ ON COMMIT { PRESERVE ROWS | DELETE ROWS | DROP } ]
// [ TABLESPACE tablespace_name ]
type Table struct {
	// scope GLOBAL | LOCAL
	scope string
	// table name
	name string
	// definitions
	definitions columnDefinitions
	// of type definition
	ofTypeDefinition ofType
	// of partition definition
	ofPartition ofPartition
	// inherits
	inherits expression
	// using method
	using string
	// Partition
	partition partitionTable
	// with
	with detailedExpression
	// tablespace
	tablespace string
	// on commit params
	onCommit string
	// is table temporary
	temp bool
	// unLogged table
	unLogged bool
	// If not exists
	ifNotExists bool
}

// IsEmpty check if table is empty
func (t *Table) IsEmpty() bool {
	return t == nil || (t.scope == "" &&
		t.name == "" &&
		len(t.definitions) == 0 &&
		t.ofTypeDefinition.IsEmpty() &&
		t.ofPartition.IsEmpty() &&
		t.inherits.Len() == 0 &&
		t.using == "" &&
		t.with.IsEmpty() &&
		t.tablespace == "" &&
		t.onCommit == "")
}

// String render table
func (t *Table) String() string {
	if t.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("CREATE")
	if t.scope != "" {
		b.WriteString(" " + t.scope)
	}
	if t.temp {
		b.WriteString(" TEMPORARY")
	} else if t.unLogged {
		b.WriteString(" UNLOGGED")
	}
	b.WriteString(" TABLE")
	if t.ifNotExists {
		b.WriteString(" IF NOT EXISTS")
	}
	if t.name != "" {
		b.WriteString(" " + t.name)
	}
	if t.definitions.Len() > 0 {
		b.WriteString(" (" + t.definitions.String() + ")")
	} else if !t.ofTypeDefinition.IsEmpty() {
		b.WriteString(t.ofTypeDefinition.String())
	} else if !t.ofPartition.IsEmpty() {
		b.WriteString(t.ofPartition.String())
	}
	if t.inherits.Len() > 0 {
		b.WriteString(" INHERITS " + t.inherits.String(", "))
	}
	if !t.partition.IsEmpty() {
		b.WriteString(" PARTITION BY " + t.partition.String())
	}
	if t.using != "" {
		b.WriteString(" USING " + t.using)
	}
	if !t.with.IsEmpty() {
		if t.with.GetDetail() != "" {
			b.WriteString(" " + t.with.GetDetail())
		} else {
			b.WriteString(" WITH " + t.with.String())
		}
	}
	if t.onCommit != "" {
		b.WriteString(" ON COMMIT " + t.onCommit)
	}
	if t.tablespace != "" {
		b.WriteString(" TABLESPACE " + t.tablespace)
	}
	return b.String() + ";"
}

// Using set using
func (t *Table) Using(using string) *Table {
	t.using = using
	return t
}

// Definitions return column definitions
func (t *Table) Definitions() *columnDefinitions {
	return &t.definitions
}

// AddColumn add column
func (t *Table) AddColumn(name string) *column {
	def, _ := t.definitions.Add()
	return def.Column().Name(name)
}

// AddForeignKey add foreign key
func (t *Table) AddForeignKey(target string, columns ...string) *foreignKey {
	fk := t.definitions.AddConstraint().ForeignKey()
	fk.Columns().Add(columns...)
	fk.References().RefTable(target)
	fk.References().Columns().Add(columns...)
	return fk
}

// AddConstraint add constraint
func (t *Table) AddConstraint() *constraintTable {
	return t.definitions.AddConstraint()
}

// With expression
func (t *Table) With(expr ...string) *detailedExpression {
	t.with.Expression().Add(expr...)
	return &t.with
}

// WithOutOIDS expression
func (t *Table) WithOutOIDS() *detailedExpression {
	t.with.SetDetail(WithWithoutOIDS)
	return &t.with
}

// Partition expression
func (t *Table) Partition() *partitionTable {
	return &t.partition
}

// OfType get of type definition
func (t *Table) OfType() *ofType {
	return &t.ofTypeDefinition
}

// OfPartition get of partition definition
func (t *Table) OfPartition() *ofPartition {
	return &t.ofPartition
}

// OnCommit set onCommit
func (t *Table) OnCommit(onCommit string) *Table {
	t.onCommit = onCommit
	return t
}

// IfNotExists Set to true
func (t *Table) IfNotExists() *Table {
	t.ifNotExists = true
	return t
}

// UnLogged Set to true
func (t *Table) UnLogged() *Table {
	t.unLogged = true
	return t
}

// Temp Set temp to true
func (t *Table) Temp() *Table {
	t.temp = true
	return t
}

// Scope set scope
func (t *Table) Scope(scope string) *Table {
	t.scope = scope
	return t
}

// Name Set name
func (t *Table) Name(name string) *Table {
	t.name = name
	return t
}

// GetName get name of table
func (t *Table) GetName() string {
	return t.name
}

// Inherits inherit form tables
func (t *Table) Inherits() *expression {
	return &t.inherits
}

// TableSpace set table space
func (t *Table) TableSpace(space string) *Table {
	t.tablespace = space
	return t
}

// SQL Render query
func (t *Table) SQL() (query string, params []any, returning []any) {
	query = t.String()
	return
}

// CreateTable init table
func CreateTable(name string) *Table {
	return &Table{name: name}
}
