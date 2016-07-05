package players

import (
	"common"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Player struct {
	PublicId string   `json:"publicId" db:"publicId"`
	Name     string   `json:"name" db:"name"`
	Age      int      `json:"age,string" db:"age"`
	Gender   string   `json:"gender" db:"gender"`
	Teams    []string `json:"teams"`
}

type playerTeamDto struct {
	PlayerId     string         `db:"playerPublicId"`
	PlayerName   string         `db:"playerName"`
	PlayerAge    int            `db:"playerAge"`
	PlayerGender string         `db:"playerGender"`
	TeamId       sql.NullString `db:"teamPublicId"`
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

func (r *repo) ListPlayers() ([]Player, error) {
	listStmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return nil, err
	}

	dtos := []playerTeamDto{}
	if err := listStmt.Select(&dtos, playerTeamDto{}); err != nil {
		return nil, err
	}

	return r.dtosToPlayers(dtos), nil
}

func (r *repo) FetchPlayer(id string) (*Player, error) {
	fetchStmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	dtos := []playerTeamDto{}
	arg := struct{ Id string }{Id: id}
	if err := fetchStmt.Select(&dtos, arg); err != nil {
		return nil, err
	}

	players := r.dtosToPlayers(dtos)
	if len(players) != 1 {
		return nil, fmt.Errorf("Received not exactly 1 player")
	}

	return &players[0], nil
}

func (r *repo) fetchPlayerPrivate(id int64, tx *sqlx.Tx) (*Player, error) {
	stmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return nil, err
	}

	if tx != nil {
		stmt = tx.NamedStmt(stmt)
	}

	dtos := []playerTeamDto{}
	arg := struct{ Id int64 }{Id: id}
	if err := stmt.Select(&dtos, arg); err != nil {
		return nil, err
	}

	players := r.dtosToPlayers(dtos)
	if len(players) != 1 {
		return nil, fmt.Errorf("Received not exactly 1 player")
	}

	return &players[0], nil
}

func (r *repo) CreatePlayer(p *Player) (*Player, error) {
	stmt, err := cache.Load(r.db, "create", createQuery)
	if err != nil {
		return nil, err
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt = tx.NamedStmt(stmt)

	res, err := stmt.Exec(p)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if player, err := r.fetchPlayerPrivate(id, tx); err != nil {
		return nil, err
	} else {
		p.PublicId = player.PublicId
	}

	for _, t := range p.Teams {
		if err = r.AddPlayerToTeam(p.PublicId, t, tx); err != nil {
			return nil, err
		}
	}

	tx.Commit()
	return p, nil
}

func (r *repo) UpdatePlayer(id string, p *Player) (*Player, error) {
	stmt, err := cache.Load(r.db, "update", updateQuery)
	if err != nil {
		return nil, err
	}

	tx, err := r.db.Beginx()
	defer tx.Rollback() // Will fail with ErrTxDone if we committed.  Meh.
	if err != nil {
		return nil, err
	}

	stmt = tx.NamedStmt(stmt)

	p.PublicId = id
	res, err := stmt.Exec(p)
	if err != nil {
		return nil, err
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if n != 1 {
		return nil, fmt.Errorf("Failed to update player.  Bad gender value?")
	}

	if p.Teams != nil {
		if err = r.clearTeamsFromPlayer(id, tx); err != nil {
			return nil, err
		}
		for _, t := range p.Teams {
			if err = r.AddPlayerToTeam(id, t, tx); err != nil {
				return nil, err
			}
		}
	}

	tx.Commit()
	return r.FetchPlayer(id)
}

func (r *repo) AddPlayerToTeam(playerId, teamId string, tx *sqlx.Tx) error {
	stmt, err := cache.Load(r.db, "addToTeam", addToTeamQuery)
	if err != nil {
		return err
	}

	if tx != nil {
		stmt = tx.NamedStmt(stmt)
	}

	arg := struct {
		PlayerId string `db:"playerId"`
		TeamId   string `db:"teamId"`
	}{
		PlayerId: playerId,
		TeamId:   teamId,
	}

	res, err := stmt.Exec(arg)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n != 1 {
		return fmt.Errorf("Affected rows != 1: %d", n)
	}

	return nil
}

func (r *repo) clearTeamsFromPlayer(playerId string, tx *sqlx.Tx) error {
	stmt, err := cache.Load(r.db, "clearTeams", clearTeamsQuery)
	if err != nil {
		return err
	}
	if tx != nil {
		stmt = tx.NamedStmt(stmt)
	}

	arg := struct{ Id string }{Id: playerId}
	_, err = stmt.Exec(arg)
	return err
}

func (r *repo) dtosToPlayers(dtos []playerTeamDto) []Player {
	players := make(map[string]Player)
	for _, d := range dtos {
		if p, ok := players[d.PlayerId]; ok {
			if d.TeamId.Valid {
				p.Teams = append(p.Teams, d.TeamId.String)
			}
		} else {
			p = Player{
				PublicId: d.PlayerId,
				Name:     d.PlayerName,
				Age:      d.PlayerAge,
				Gender:   d.PlayerGender,
				Teams:    []string{},
			}

			if d.TeamId.Valid {
				p.Teams = []string{d.TeamId.String}
			}

			players[d.PlayerId] = p
		}
	}

	res := make([]Player, 0, len(players))
	for _, v := range players {
		res = append(res, v)
	}

	return res
}
