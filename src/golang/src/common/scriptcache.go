package common

import (
	"github.com/jmoiron/sqlx"
)

type scriptCache struct {
	statements map[*sqlx.DB]map[string]*sqlx.NamedStmt
}

func CreateScriptCache() *scriptCache {
	return &scriptCache{
		statements: make(map[*sqlx.DB]map[string]*sqlx.NamedStmt),
	}
}

func (s *scriptCache) Load(db *sqlx.DB, name, script string) (stmt *sqlx.NamedStmt, err error) {
	var ok bool
	if _, ok = s.statements[db]; !ok {
		s.statements[db] = make(map[string]*sqlx.NamedStmt)
	}

	if stmt, ok = s.statements[db][name]; !ok {
		stmt, err = db.PrepareNamed(script)
	} else {
		err = nil
	}

	return
}
