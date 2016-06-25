package leagues

import (
	"common"
	"divisions"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type League struct {
	PublicId    string `json:"publicId" db:"publicId"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Location    string `json:"location" db:"location"`
	Division    string `json:"division" db:"division"`
	Gender      string `json:"gender" db:"gender"`
	StartDate   string `json:"startDate" db:"startDate"`
	EndDate     string `json:"endDate" db:"endDate"`
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

func (r *repo) ListLeagues() ([]League, error) {
	listStmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return nil, err
	}

	leagues := []League{}
	if err := listStmt.Select(&leagues, League{}); err != nil {
		return nil, err
	}
	return leagues, nil
}

func (r *repo) FetchLeague(id string) (*League, error) {
	fetchStmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	league := League{}
	arg := struct{ Id string }{Id: id}
	if err := fetchStmt.Get(&league, arg); err != nil {
		return nil, err
	}
	return &league, nil
}

func (r *repo) fetchLeaguePrivate(id int) (*League, error) {
	fetchPrivateStmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return nil, err
	}

	league := League{}
	arg := struct{ Id int }{Id: id}
	if err := fetchPrivateStmt.Get(&league, &arg); err != nil {
		return nil, err
	}
	return &league, nil
}

func (r *repo) CreateLeague(l *League) (*League, error) {
	divisionRepo := divisions.CreateRepo(r.db)
	_, err := divisionRepo.CreateIfNotExists(l.Division)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Division: %s", err)
	}

	createStmt, err := cache.Load(r.db, "create", createQuery)
	if err != nil {
		return nil, err
	}

	res, err := createStmt.Exec(l)
	if err != nil {
		return nil, err
	}

	if id, err := res.LastInsertId(); err != nil {
		return nil, err
	} else {
		return r.fetchLeaguePrivate(int(id))
	}
}

func (r *repo) ReplaceLeague(id string, p *League) (*League, error) {
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
		return nil, fmt.Errorf("Failed to update league.  Bad gender value?")
	}

	return r.FetchLeague(id)
}

func idStruct(id int) interface{} {
	return struct {
		Id int `db:"id"`
	}{
		Id: id,
	}
}
