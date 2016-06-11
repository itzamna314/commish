package matches

import (
	"common"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Match struct {
	PublicId   string `json:"publicId" db:"publicId"`
	HomeTeamId string `json:"homeTeam" db:"homeTeamId"`
	AwayTeamId string `json:"awayTeam" db:"awayTeamId"`
	State      string `json:"state" db:"state"`
}

type repo struct {
	db *sqlx.DB
}

var (
	cache = common.CreateScriptCache()
)

func createRepo(db *sqlx.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) ListMatches() (matches []Match, err error) {
	matches = []Match{}
	stmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return
	}

	err = stmt.Select(&matches, Match{})
	return
}

func (r *repo) FetchMatch(publicId string) (match *Match, err error) {
	match = &Match{}
	stmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return
	}

	arg := struct{ Id string }{Id: publicId}
	err = stmt.Get(match, arg)
	return
}

func (r *repo) fetchMatchPrivate(id int64) (match *Match, err error) {
	match = &Match{}
	stmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return
	}

	arg := struct{ Id int64 }{Id: id}
	err = stmt.Get(match, arg)
	fmt.Printf("Id: %d, Fetched match: %+v\n", id, match)
	return
}

func (r *repo) CreateMatch(m *Match) (match *Match, err error) {
	match = &Match{}
	stmt, err := cache.Load(r.db, "create", createQuery)
	if err != nil {
		return
	}

	res, err := stmt.Exec(m)
	if err != nil {
		return
	}

	if n, raErr := res.RowsAffected(); raErr == nil {
		if n != 1 {
			err = fmt.Errorf("Failed to create match")
			return
		}
	}

	if id, err := res.LastInsertId(); err == nil {
		match, err = r.fetchMatchPrivate(id)
	}

	return
}

func (r *repo) ReplaceMatch(id string, m *Match) (match *Match, err error) {
	match = &Match{}
	stmt, err := cache.Load(r.db, "replace", replaceQuery)
	if err != nil {
		return
	}

	m.PublicId = id
	res, err := stmt.Exec(m)
	if err != nil {
		return
	}

	if n, err := res.RowsAffected(); err == nil {
		if n == 1 {
			match, err = r.FetchMatch(id)
		} else {
			err = fmt.Errorf("Failed to update match.")
		}
	}
	return
}
