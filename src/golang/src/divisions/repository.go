package divisions

import (
	"common"
	"github.com/jmoiron/sqlx"
)

type Division struct {
	PublicId string `json:"publicId" db:"publicId"`
	Name     string `json:"name" db:"name"`
}

type repo struct {
	db *sqlx.DB
}

var (
	cache = common.CreateScriptCache()
)

func CreateRepo(db *sqlx.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateIfNotExists(name string, tx *sqlx.Tx) (*Division, error) {
	fetchStmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	if tx != nil {
		fetchStmt = tx.NamedStmt(fetchStmt)
	}

	arg := struct {
		Name string `db:"name"`
	}{Name: name}

	division := Division{}
	divisions := []Division{}
	if err = fetchStmt.Select(&divisions, arg); err != nil {
		return nil, err
	}

	if len(divisions) != 1 {
		createStmt, err := cache.Load(r.db, "create", createQuery)
		if err != nil {
			return nil, err
		}

		if tx != nil {
			createStmt = tx.NamedStmt(createStmt)
		}

		_, err = createStmt.Exec(arg)
		if err != nil {
			return nil, err
		}

		if err = fetchStmt.Get(&division, arg); err != nil {
			return nil, err
		}
	}

	return &division, nil
}
