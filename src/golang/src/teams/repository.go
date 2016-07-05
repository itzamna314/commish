package teams

import (
	"common"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Team struct {
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

func (r *repo) ListTeams() ([]Team, error) {
	listStmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return nil, err
	}

	teams := []Team{}
	if err := listStmt.Select(&teams, Team{}); err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *repo) FetchTeam(id string) (*Team, error) {
	fetchStmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	team := Team{}
	arg := struct{ Id string }{Id: id}
	if err := fetchStmt.Get(&team, arg); err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *repo) fetchTeamPrivate(id int) (*Team, error) {
	fetchPrivateStmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return nil, err
	}

	team := Team{}
	arg := struct{ Id int }{Id: id}
	if err := fetchPrivateStmt.Get(&team, &arg); err != nil {
		return nil, err
	}
	return &team, nil
}

func (r *repo) CreateTeam(p *Team) (*Team, error) {
	createStmt, err := cache.Load(r.db, "create", createQuery)
	if err != nil {
		return nil, err
	}

	res, err := createStmt.Exec(p)
	if err != nil {
		return nil, err
	}

	if id, err := res.LastInsertId(); err != nil {
		return nil, err
	} else {
		return r.fetchTeamPrivate(int(id))
	}
}

func (r *repo) ReplaceTeam(id string, p *Team) (*Team, error) {
	replaceStmt, err := cache.Load(r.db, "replace", replaceQuery)
	if err != nil {
		return nil, err
	}

	p.PublicId = id
	res, err := replaceStmt.Exec(p)
	if err != nil {
		return nil, err
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if n != 1 {
		return nil, fmt.Errorf("Failed to update team.  Bad gender value?")
	}

	return r.FetchTeam(id)
}

func idStruct(id int) interface{} {
	return struct {
		Id int `db:"id"`
	}{
		Id: id,
	}
}
