package gosql

import (
	"strings"
)

const (
	// OwnerCurrentRole role
	OwnerCurrentRole = "CURRENT_ROLE"
	// OwnerCurrentUser user
	OwnerCurrentUser = "CURRENT_USER"
	// OwnerCurrentSession session
	OwnerCurrentSession = "SESSION_USER"
)

// where action is one of:
//
// ADD [ COLUMN ] [ IF NOT EXISTS ] column_name data_type [ COLLATE collation ] [ column_constraint [ ... ] ]
// DROP [ COLUMN ] [ IF EXISTS ] column_name [ RESTRICT | CASCADE ]
// ALTER [ COLUMN ] column_name [ SET DATA ] TYPE data_type [ COLLATE collation ] [ USING expression ]
// ALTER [ COLUMN ] column_name SET DEFAULT expression
// ALTER [ COLUMN ] column_name DROP DEFAULT
// ALTER [ COLUMN ] column_name { SET | DROP } NOT NULL
// ALTER [ COLUMN ] column_name DROP EXPRESSION [ IF EXISTS ]
// ALTER [ COLUMN ] column_name ADD GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence_options ) ]
// ALTER [ COLUMN ] column_name { SET GENERATED { ALWAYS | BY DEFAULT } | SET sequence_option | RESTART [ [ WITH ] restart ] } [...]
// ALTER [ COLUMN ] column_name DROP IDENTITY [ IF EXISTS ]
// ALTER [ COLUMN ] column_name SET STATISTICS integer
// ALTER [ COLUMN ] column_name SET ( attribute_option = value [, ... ] )
// ALTER [ COLUMN ] column_name RESET ( attribute_option [, ... ] )
// ALTER [ COLUMN ] column_name SET STORAGE { PLAIN | EXTERNAL | EXTENDED | MAIN }
// ALTER [ COLUMN ] column_name SET COMPRESSION compression_method
// ADD table_constraint [ NOT VALID ]
// ADD table_constraint_using_index
// ALTER CONSTRAINT constraint_name [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
// VALIDATE CONSTRAINT constraint_name
// DROP CONSTRAINT [ IF EXISTS ]  constraint_name [ RESTRICT | CASCADE ]
// DISABLE TRIGGER [ trigger_name | ALL | USER ]
// ENABLE TRIGGER [ trigger_name | ALL | USER ]
// ENABLE REPLICA TRIGGER trigger_name
// ENABLE ALWAYS TRIGGER trigger_name
// DISABLE RULE rewrite_rule_name
// ENABLE RULE rewrite_rule_name
// ENABLE REPLICA RULE rewrite_rule_name
// ENABLE ALWAYS RULE rewrite_rule_name
// DISABLE ROW LEVEL SECURITY
// ENABLE ROW LEVEL SECURITY
// FORCE ROW LEVEL SECURITY
// NO FORCE ROW LEVEL SECURITY
// CLUSTER ON index_name
// SET WITHOUT CLUSTER
// SET WITHOUT OIDS
// SET ACCESS METHOD new_access_method
// SET TABLESPACE new_tablespace
// SET { LOGGED | UNLOGGED }
// SET ( storage_parameter [= value] [, ... ] )
// RESET ( storage_parameter [, ... ] )
// INHERIT parent_table
// NO INHERIT parent_table
// OF type_name
// NOT OF
// OWNER TO { new_owner | CURRENT_ROLE | CURRENT_USER | SESSION_USER }
// REPLICA IDENTITY { DEFAULT | USING INDEX index_name | FULL | NOTHING }
type alterTableAction struct {
	// common
	ordered orderedExpression
	// add action
	add alterTableActionAdd
	// set action
	set alterTableActionSet
	// drop action
	drop alterTableActionDrop
}

// AlterColumn alter column
func (a *alterTableAction) AlterColumn(name string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ALTER COLUMN ", name))
	return a
}

// Constraint alter constraint
func (a *alterTableAction) Constraint(name string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ALTER CONSTRAINT ", name))
	return a
}

// Deferrable set deferrable
func (a *alterTableAction) Deferrable() *alterTableAction {
	a.ordered.Add(1, a.ordered.Concat("DEFERRABLE"))
	return a
}

// NotDeferrable set not deferrable
func (a *alterTableAction) NotDeferrable() *alterTableAction {
	a.ordered.Add(1, a.ordered.Concat("NOT DEFERRABLE"))
	return a
}

// InitiallyDeferred set INITIALLY DEFERRED
func (a *alterTableAction) InitiallyDeferred() *alterTableAction {
	a.ordered.Add(1, a.ordered.Concat("INITIALLY DEFERRED"))
	return a
}

// InitiallyImmediate set INITIALLY IMMEDIATE
func (a *alterTableAction) InitiallyImmediate() *alterTableAction {
	a.ordered.Add(1, a.ordered.Concat("INITIALLY IMMEDIATE"))
	return a
}

// ValidateConstraint validate constraint
func (a *alterTableAction) ValidateConstraint(constraint string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("VALIDATE CONSTRAINT ", constraint))
	return a
}

// DisableTrigger disable trigger
func (a *alterTableAction) DisableTrigger(trigger string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("DISABLE TRIGGER ", trigger))
	return a
}

// DisableRule disable rule
func (a *alterTableAction) DisableRule(rule string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("DISABLE RULE ", rule))
	return a
}

// DisableRowLevelSecurity disable row level security
func (a *alterTableAction) DisableRowLevelSecurity() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("DISABLE ROW LEVEL SECURITY"))
	return a
}

// EnableTrigger enable trigger
func (a *alterTableAction) EnableTrigger(trigger string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE TRIGGER ", trigger))
	return a
}

// EnableReplicaTrigger enable replica trigger
func (a *alterTableAction) EnableReplicaTrigger(trigger string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE REPLICA TRIGGER ", trigger))
	return a
}

// EnableAlwaysTrigger enable always trigger
func (a *alterTableAction) EnableAlwaysTrigger(trigger string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE ALWAYS TRIGGER ", trigger))
	return a
}

// EnableRule enable rule
func (a *alterTableAction) EnableRule(rule string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE RULE ", rule))
	return a
}

// EnableReplicaRule enable replica rule
func (a *alterTableAction) EnableReplicaRule(rule string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE REPLICA RULE ", rule))
	return a
}

// EnableAlwaysRule enable always rule
func (a *alterTableAction) EnableAlwaysRule(rule string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE ALWAYS RULE ", rule))
	return a
}

// EnableRowLevelSecurity enable row level security
func (a *alterTableAction) EnableRowLevelSecurity() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("ENABLE ROW LEVEL SECURITY"))
	return a
}

// ForceRowLevelSecurity force row level security
func (a *alterTableAction) ForceRowLevelSecurity() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("FORCE ROW LEVEL SECURITY"))
	return a
}

// NoForceRowLevelSecurity no force row level security
func (a *alterTableAction) NoForceRowLevelSecurity() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("NO FORCE ROW LEVEL SECURITY"))
	return a
}

// ClusterOn cluster on index
func (a *alterTableAction) ClusterOn(index string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("CLUSTER ON ", index))
	return a
}

// ResetOption reset options
func (a *alterTableAction) ResetOption(option ...string) *alterTableAction {
	a.ordered.Delimiter(", ")
	if a.ordered.Iterator() == 0 {
		a.ordered.iterator = 1
		a.ordered.Append(a.ordered.Concat("RESET ("))
	}
	for i := range option {
		a.ordered.Append(a.ordered.Concat(option[i]))
	}
	return a
}

// ResetStorageParams reset storage params
func (a *alterTableAction) ResetStorageParams(parameter ...string) *alterTableAction {
	a.ordered.Delimiter(", ")
	if a.ordered.Iterator() == 0 {
		a.ordered.Append(a.ordered.Concat("RESET ("))
	}
	for i := range parameter {
		a.ordered.Append(a.ordered.Concat(parameter[i]))
	}
	return a
}

// Inherit parent table
func (a *alterTableAction) Inherit(table string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("INHERIT ", table))
	return a
}

// NoInherit parent table
func (a *alterTableAction) NoInherit(table string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("NO INHERIT ", table))
	return a
}

// Of type name
func (a *alterTableAction) Of(typeName string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("OF ", typeName))
	return a
}

// NotOf not of
func (a *alterTableAction) NotOf() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("NOT OF"))
	return a
}

// OwnerTo set new owner
func (a *alterTableAction) OwnerTo(owner string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("OWNER TO ", owner))
	return a
}

// ReplicaIdentityDefault set default replica identity
func (a *alterTableAction) ReplicaIdentityDefault() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("REPLICA IDENTITY DEFAULT"))
	return a
}

// ReplicaIdentityUsingIndex set replica identity using index
func (a *alterTableAction) ReplicaIdentityUsingIndex(index string) *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("REPLICA IDENTITY USING INDEX ", index))
	return a
}

// ReplicaIdentityFull set replica identity full
func (a *alterTableAction) ReplicaIdentityFull() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("REPLICA IDENTITY FULL"))
	return a
}

// ReplicaIdentityNothing set replica identity nothing
func (a *alterTableAction) ReplicaIdentityNothing() *alterTableAction {
	a.ordered.Add(0, a.ordered.Concat("REPLICA IDENTITY NOTHING"))
	return a
}

// Add action
func (a *alterTableAction) Add() *alterTableActionAdd {
	a.set.Reset()
	a.drop.Reset()
	return &a.add
}

// Set action
func (a *alterTableAction) Set() *alterTableActionSet {
	a.add.Reset()
	a.drop.Reset()
	return &a.set
}

// Drop action
func (a *alterTableAction) Drop() *alterTableActionDrop {
	a.set.Reset()
	a.drop.Reset()
	return &a.drop
}

// IsEmpty check if empty
func (a *alterTableAction) IsEmpty() bool {
	return a == nil || a.ordered.IsEmpty() && a.add.IsEmpty() && a.set.IsEmpty() && a.drop.IsEmpty()
}

// Reset reset item data
func (a *alterTableAction) Reset() *alterTableAction {
	a.ordered.Reset()
	a.add.Reset()
	a.set.Reset()
	a.drop.Reset()
	return a
}

// Grow memory
func (a *alterTableAction) Grow(n int) *alterTableAction {
	a.ordered.Grow(n)
	return a
}

// GetArguments get arguments
func (a *alterTableAction) GetArguments() []any {
	return append(append(append(a.ordered.GetArguments(), a.add.GetArguments()...), a.set.GetArguments()...), a.drop.GetArguments()...)
}

// String render alter table query
func (a *alterTableAction) String() string {
	if a.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if !a.ordered.IsEmpty() {
		if a.ordered.Iterator() > 0 {
			b.WriteString(a.ordered.String() + ")")
		} else {
			b.WriteString(a.ordered.String())
		}
		if !a.add.IsEmpty() || !a.set.IsEmpty() || !a.drop.IsEmpty() {
			b.WriteString(" ")
		}
	}
	if !a.add.IsEmpty() {
		b.WriteString(a.add.String())
	} else if !a.set.IsEmpty() {
		b.WriteString(a.set.String())
	} else if !a.drop.IsEmpty() {
		b.WriteString(a.drop.String())
	}
	return b.String()
}
