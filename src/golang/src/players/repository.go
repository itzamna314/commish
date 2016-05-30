package players

import (
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
	listStmt         *sqlx.NamedStmt
	fetchPrivateStmt *sqlx.NamedStmt
	createStmt       *sqlx.NamedStmt
)

func createRepo(db *sqlx.DB) *repo {
	// Ideally, we would only init these
	// prepared statements once.
	var err error
	listStmt, err = db.PrepareNamed(listQuery)
	fetchPrivateStmt, err = db.PrepareNamed(fetchPrivateQuery)
	createStmt, err = db.PrepareNamed(createQuery)
	if err != nil {
		panic(err)
	}

	return &repo{
		db: db,
	}
}

func (r *repo) ListPlayers() ([]Player, error) {
	players := []Player{}
	if err := listStmt.Select(&players, Player{}); err != nil {
		return nil, err
	}
	return players, nil
}

func (r *repo) FetchPlayer(id string) (*Player, error) {
	player := Player{}
	if err := listStmt.Select(&player, id); err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *repo) fetchPlayerPrivate(id int) (*Player, error) {
	player := Player{}
	arg := struct{ Id int }{Id: id}
	if err := fetchPrivateStmt.Get(&player, &arg); err != nil {
		return nil, err
	}
	return &player, nil
}

func (r *repo) CreatePlayer(p *Player) (int, error) {
	res, err := createStmt.Exec(p)
	if err != nil {
		return 0, err
	}

	if id, err := res.LastInsertId(); err != nil {
		return 0, err
	} else {
		return int(id), nil
	}
}

func idStruct(id int) interface{} {
	return struct {
		Id int `db:"id"`
	}{
		Id: id,
	}
}
