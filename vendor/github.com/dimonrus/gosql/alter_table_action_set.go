package gosql

import (
	"strconv"
)

type Storage string

const (
	// StoragePlain PLAIN
	StoragePlain Storage = "PLAIN"
	// StorageExternal EXTERNAL
	StorageExternal Storage = "EXTERNAL"
	// StorageExtended EXTENDED
	StorageExtended Storage = "EXTENDED"
	// StorageMain MAIN
	StorageMain Storage = "MAIN"
	// StorageDefault DEFAULT
	StorageDefault Storage = "DEFAULT"
)

// where set action is one of:
//
// [ SET DATA ] TYPE data_type [ COLLATE collation ] [ USING expression ]
// SET DEFAULT expression
// SET NOT NULL
// { SET GENERATED { ALWAYS | BY DEFAULT } | SET sequence_option | RESTART [ [ WITH ] restart ] } [...]
// SET STATISTICS integer
// SET ( attribute_option = value [, ... ] )
// SET STORAGE { PLAIN | EXTERNAL | EXTENDED | MAIN }
// SET COMPRESSION compression_method
// SET WITHOUT CLUSTER
// SET WITHOUT OIDS
// SET ACCESS METHOD new_access_method
// SET SCHEMA new_schema
// SET TABLESPACE new_tablespace
// SET { LOGGED | UNLOGGED }
// SET ( storage_parameter [= value] [, ... ] )
type alterTableActionSet struct {
	// ordered expression
	ordered orderedExpression
}

// DataType set data type
func (a *alterTableActionSet) DataType(dataType string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET DATA TYPE ", dataType))
	return a
}

// Collate set collation
func (a *alterTableActionSet) Collate(collation string) *alterTableActionSet {
	a.ordered.Add(1, a.ordered.Concat("COLLATE ", collation))
	return a
}

// Using set using
func (a *alterTableActionSet) Using(expression string) *alterTableActionSet {
	a.ordered.Add(2, a.ordered.Concat("USING ", expression))
	return a
}

// Default set default
func (a *alterTableActionSet) Default(expression string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET DEFAULT ", expression))
	return a
}

// NotNull set not null
func (a *alterTableActionSet) NotNull() *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET NOT NULL"))
	return a
}

// GeneratedAlways set generated always
func (a *alterTableActionSet) GeneratedAlways(order int) *alterTableActionSet {
	a.ordered.Add(order, a.ordered.Concat("SET GENERATED ALWAYS"))
	return a
}

// GeneratedByDefault set generated by default
func (a *alterTableActionSet) GeneratedByDefault(order int) *alterTableActionSet {
	a.ordered.Add(order, a.ordered.Concat("SET GENERATED BY DEFAULT"))
	return a
}

// Sequence set sequence option
func (a *alterTableActionSet) Sequence(order int, option string) *alterTableActionSet {
	a.ordered.Add(order, a.ordered.Concat("SET ", option))
	return a
}

// Restart restart
func (a *alterTableActionSet) Restart(order int, restart string) *alterTableActionSet {
	a.ordered.Add(order, a.ordered.Concat("RESTART ", restart))
	return a
}

// Statistics Set statistics
func (a *alterTableActionSet) Statistics(value int) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET STATISTICS ", strconv.FormatInt(int64(value), 10)))
	return a
}

// AttributeOptions Set attribute option
func (a *alterTableActionSet) AttributeOptions(option ...string) *alterTableActionSet {
	if len(option) == 0 {
		return a
	}
	a.ordered.Delimiter(", ")
	if a.ordered.Iterator() == 0 {
		a.ordered.Append(a.ordered.Concat("SET (", option[0]))
		for i := range option[1:] {
			a.ordered.Append(a.ordered.Concat(option[1:][i]))
		}
	} else {
		for i := range option {
			a.ordered.Append(a.ordered.Concat(option[i]))
		}
	}
	return a
}

// Storage Set storage
func (a *alterTableActionSet) Storage(storage Storage) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET ", string(storage)))
	return a
}

// Compression Set compression method
func (a *alterTableActionSet) Compression(method string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET COMPRESSION ", method))
	return a
}

// WithoutCluster Set without cluster
func (a *alterTableActionSet) WithoutCluster() *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET WITHOUT CLUSTER"))
	return a
}

// WithoutOIDS Set without oids
func (a *alterTableActionSet) WithoutOIDS() *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET WITHOUT OIDS"))
	return a
}

// AccessMethod Set access method
func (a *alterTableActionSet) AccessMethod(method string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET ACCESS METHOD ", method))
	return a
}

// Schema Set new schema
func (a *alterTableActionSet) Schema(schema string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET SCHEMA ", schema))
	return a
}

// TableSpace Set table space
func (a *alterTableActionSet) TableSpace(tableSpace string) *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET TABLESPACE ", tableSpace))
	return a
}

// Logged Set logged
func (a *alterTableActionSet) Logged() *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET LOGGED"))
	return a
}

// UnLogged Set unlogged
func (a *alterTableActionSet) UnLogged() *alterTableActionSet {
	a.ordered.Add(0, a.ordered.Concat("SET UNLOGGED"))
	return a
}

// StorageParameters Set storage parameter
func (a *alterTableActionSet) StorageParameters(param string, args ...any) *alterTableActionSet {
	a.ordered.Delimiter(", ")
	if a.ordered.Iterator() == 0 {
		a.ordered.Append(a.ordered.Concat("SET (", param), args...)
	} else {
		a.ordered.Append(a.ordered.Concat(param), args...)
	}
	return a
}

// IsEmpty check if empty
func (a *alterTableActionSet) IsEmpty() bool {
	return a == nil || a.ordered.IsEmpty()
}

// Reset reset item data
func (a *alterTableActionSet) Reset() *alterTableActionSet {
	a.ordered.Reset()
	return a
}

// Grow memory
func (a *alterTableActionSet) Grow(n int) *alterTableActionSet {
	a.ordered.Grow(n)
	return a
}

// GetArguments get arguments
func (a *alterTableActionSet) GetArguments() []any {
	return a.ordered.GetArguments()
}

// String render alter table query
func (a *alterTableActionSet) String() string {
	if a.IsEmpty() {
		return ""
	}
	if a.ordered.iterator > 0 {
		return a.ordered.String() + ")"
	}
	return a.ordered.String()
}