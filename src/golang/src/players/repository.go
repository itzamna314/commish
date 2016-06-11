package players

import (
	"common"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Player struct {
	PublicId string `json:"publicId" db:"publicId"`
	Name     string `json:"name" db:"name"`
	Age      int    `json:"age" db:"age"`
	Gender   string `json:"gender" db:"gender"`
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

func (r *repo) ListPlayers() ([]Player, error) {
	listStmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return nil, err
	}

	players := []Player{}
	if err := listStmt.Select(&players, Player{}); err != nil {
		return nil, err
	}
	return players, nil
}

func (r *repo) FetchPlayer(id string) (*Player, error) {
	fetchStmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	player := Player{}
	arg := struct{ Id string }{Id: id}
	if err := fetchStmt.Get(&player, arg); err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *repo) fetchPlayerPrivate(id int64) (*Player, error) {
	fetchPrivateStmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return nil, err
	}

	player := Player{}
	arg := struct{ Id int64 }{Id: id}
	if err := fetchPrivateStmt.Get(&player, &arg); err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *repo) CreatePlayer(p *Player) (*Player, error) {
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
		return r.fetchPlayerPrivate(id)
	}
}

func (r *repo) ReplacePlayer(id string, p *Player) (*Player, error) {
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
		return nil, fmt.Errorf("Failed to update player.  Bad gender value?")
	}

	return r.FetchPlayer(id)
}
