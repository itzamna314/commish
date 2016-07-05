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

type PlayerTeamDto struct {
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
	fmt.Printf("Listing players...\n")
	listStmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		fmt.Printf("Failed to load prepared statement\n")
		return nil, err
	}

	fmt.Printf("Loaded prepared statement...\n")

	rows, err := r.db.Query(listQuery)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Queried db\n")

	for rows.Next() {
		dto := PlayerTeamDto{}
		rows.Scan(&dto.PlayerId, &dto.PlayerName, &dto.PlayerAge, &dto.PlayerGender, &dto.TeamId)
		fmt.Printf("%+v\n", dto)
	}

	dtos := []PlayerTeamDto{}
	if err := listStmt.Select(&dtos, PlayerTeamDto{}); err != nil {
		return nil, err
	}

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

	return res, nil
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

func (r *repo) UpdatePlayer(id string, p *Player) (*Player, error) {
	updateStmt, err := cache.Load(r.db, "update", updateQuery)
	if err != nil {
		return nil, err
	}

	p.PublicId = id
	res, err := updateStmt.Exec(p)
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
