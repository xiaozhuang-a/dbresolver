package dbresolver

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Operation specifies dbresolver mode
type Operation string

const (
	writeName = "gorm:db_resolver:write"
	readName  = "gorm:db_resolver:read"
)

// ModifyStatement modify operation mode
func (op Operation) ModifyStatement(stmt *gorm.Statement) {
	var optName string
	if op == Write {
		optName = writeName
	} else if op == Read {
		optName = readName
	}

	if optName != "" {
		stmt.Clauses[optName] = clause.Clause{}
		stmt.DB.Callback().Query().Get("gorm:db_resolver")(stmt.DB)
	}
}

// Build implements clause.Expression interface
func (op Operation) Build(clause.Builder) {
}

// Use specifies configuration
func Use(str string) clause.Expression {
	return using{Use: str}
}

type using struct {
	Use string
}

const usingName = "gorm:db_resolver:using"

// ModifyStatement modify operation mode
func (u using) ModifyStatement(stmt *gorm.Statement) {
	stmt.Clauses[usingName] = clause.Clause{Expression: u}
	stmt.DB.Callback().Query().Get("gorm:db_resolver")(stmt.DB)
}

// Build implements clause.Expression interface
func (u using) Build(clause.Builder) {
}
