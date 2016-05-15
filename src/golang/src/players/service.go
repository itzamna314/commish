package players

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Player struct {
	PublicId string `json:"publicId"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

type PlayerService struct {
	ConnectionString string
	db               *sql.DB
}

func (p *PlayerService) ensureConnected() error {
	if p.db != nil {
		return nil
	}

	db, err := sql.Open("mysql", p.ConnectionString)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PlayerService) ListPlayers() ([]Player, error) {
	if err := p.ensureConnected(); err != nil {
		return nil, err
	}

	rows, err := p.db.Query(listPlayersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player

	for rows.Next() {
		p := Player{}
		if err = rows.Scan(&p.PublicId, &p.Name, &p.Age, &p.Gender); err != nil {
			return nil, err
		}

		players = append(players, p)
	}

	return players, nil
}
