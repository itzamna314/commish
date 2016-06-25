package games

import (
	"common"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Game struct {
	PublicId      string `json:"publicId" db:"publicId"`
	MatchId       string `json:"match" db:"matchId"`
	State         string `json:"state" db:"state"`
	HomeTeamScore int    `json:"homeTeamScore" db:"homeTeamScore"`
	AwayTeamScore int    `json:"homeTeamScore" db:"homeTeamScore"`
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

func (r *repo) ListGames() (games []Game, err error) {
	games = []Game{}
	stmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return
	}

	err = stmt.Select(&games, Game{})
	return
}

func (r *repo) FetchGame(publicId string) (match *Game, err error) {
	match = &Game{}
	stmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return
	}

	arg := struct{ Id string }{Id: publicId}
	err = stmt.Get(match, arg)
	return
}

func (r *repo) fetchGamePrivate(id int64) (match *Game, err error) {
	match = &Game{}
	stmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return
	}

	arg := struct{ Id int64 }{Id: id}
	err = stmt.Get(match, arg)
	fmt.Printf("Id: %d, Fetched match: %+v\n", id, match)
	return
}

func (r *repo) CreateGame(m *Game) (match *Game, err error) {
	match = &Game{}
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
		match, err = r.fetchGamePrivate(id)
	}

	return
}

func (r *repo) ReplaceGame(id string, m *Game) (match *Game, err error) {
	match = &Game{}
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
			match, err = r.FetchGame(id)
		} else {
			err = fmt.Errorf("Failed to update match.")
		}
	}
	return
}
