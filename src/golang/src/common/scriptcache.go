package common

import (
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

const FlagsToken string = "$FLAGS$"

type scriptCache struct {
	statements    map[*sqlx.DB]map[string]*sqlx.NamedStmt
	preprocessors []Preprocessor
}

type Preprocessor func(string) (string, error)

func CreateScriptCache() *scriptCache {
	return &scriptCache{
		statements:    make(map[*sqlx.DB]map[string]*sqlx.NamedStmt),
		preprocessors: make([]Preprocessor, 0, 1),
	}
}

func (s *scriptCache) AddPreprocessor(f Preprocessor) {
	s.preprocessors = append(s.preprocessors, f)
}

func (s *scriptCache) Load(db *sqlx.DB, name, script string) (stmt *sqlx.NamedStmt, err error) {
	stmt, err = s.LoadWithFlags(db, name, script, -1)
	return
}

func (s *scriptCache) LoadWithFlags(db *sqlx.DB, name, script string, flags int) (stmt *sqlx.NamedStmt, err error) {
	var ok bool
	err = nil

	if _, ok = s.statements[db]; !ok {
		s.statements[db] = make(map[string]*sqlx.NamedStmt)
	}

	if flags >= 0 {
		script = strings.Replace(script, FlagsToken, strconv.Itoa(flags), -1)
		name = name + strconv.Itoa(flags)
	}

	if stmt, ok = s.statements[db][name]; !ok {
		if stmt, err = db.PrepareNamed(script); err != nil {
			return nil, err
		}

		s.statements[db][name] = stmt
	}
	return
}
