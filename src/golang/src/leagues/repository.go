package leagues

import (
	"common"
	"database/sql"
	"divisions"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type League struct {
	PublicId    string   `json:"publicId" db:"publicId"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Location    string   `json:"location" db:"location"`
	Division    string   `json:"division" db:"division"`
	Gender      string   `json:"gender" db:"gender"`
	StartDate   string   `json:"startDate" db:"startDate"`
	EndDate     string   `json:"endDate" db:"endDate"`
	Teams       []string `json:"teams"`
}

type leagueTeamDto struct {
	League
	TeamId sql.NullString `db:"teamPublicId"`
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
	stmt, err := cache.Load(r.db, "list", listQuery)
	if err != nil {
		return nil, err
	}

	dtos := []leagueTeamDto{}
	if err := stmt.Select(&dtos, League{}); err != nil {
		return nil, err
	}

	return r.dtosToLeagues(dtos), nil
}

func (r *repo) FetchLeague(id string) (*League, error) {
	stmt, err := cache.Load(r.db, "fetch", fetchQuery)
	if err != nil {
		return nil, err
	}

	dtos := []leagueTeamDto{}
	arg := struct{ Id string }{Id: id}
	if err := stmt.Select(&dtos, arg); err != nil {
		return nil, err
	}

	leagues := r.dtosToLeagues(dtos)
	if len(leagues) != 1 {
		return nil, fmt.Errorf("Received not exactly 1 league")
	}

	return &leagues[0], nil
}

func (r *repo) fetchLeaguePrivate(id int) (*League, error) {
	stmt, err := cache.Load(r.db, "fetchPrivate", fetchPrivateQuery)
	if err != nil {
		return nil, err
	}

	dtos := []leagueTeamDto{}
	arg := struct{ Id int }{Id: id}
	if err := stmt.Select(&dtos, arg); err != nil {
		return nil, err
	}

	leagues := r.dtosToLeagues(dtos)
	if len(leagues) != 1 {
		return nil, fmt.Errorf("Received not exactly 1 league")
	}

	return &leagues[0], nil
}

func (r *repo) CreateLeague(l *League) (*League, error) {
	divisionRepo := divisions.CreateRepo(r.db)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = divisionRepo.CreateIfNotExists(l.Division, tx)
	if err != nil {
		return nil, fmt.Errorf("Failed to create division: %s", err)
	}

	stmt, err := cache.Load(r.db, "create", createQuery)
	if err != nil {
		return nil, err
	}

	stmt = tx.NamedStmt(stmt)

	res, err := stmt.Exec(l)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	if id, err := res.LastInsertId(); err != nil {
		return nil, err
	} else {
		return r.fetchLeaguePrivate(int(id))
	}
}

func (r *repo) UpdateLeague(id string, l *League) (*League, error) {
	divisionRepo := divisions.CreateRepo(r.db)
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = divisionRepo.CreateIfNotExists(l.Division, tx)
	if err != nil {
		return nil, err
	}

	stmt, err := cache.Load(r.db, "update", updateQuery)
	if err != nil {
		return nil, err
	}

	stmt = tx.NamedStmt(stmt)

	l.PublicId = id
	res, err := stmt.Exec(l)
	if err != nil {
		return nil, err
	}

	if n, err := res.RowsAffected(); err != nil {
		return nil, err
	} else if n != 1 {
		return nil, fmt.Errorf("expected 1 row affected, found %d", n)
	}

	if l.Teams != nil {
		fmt.Printf("Teams: %v\n", l.Teams)
		if err = r.clearTeamsFromLeague(id, tx); err != nil {
			return nil, err
		}
		for _, t := range l.Teams {
			if err = r.AddTeamToLeague(id, t, tx); err != nil {
				return nil, err
			}
		}
	}

	tx.Commit()
	return r.FetchLeague(id)
}

func (r *repo) AddTeamToLeague(leagueId, teamId string, tx *sqlx.Tx) error {
	if tx == nil {
		return fmt.Errorf("tx is required")
	}

	stmt, err := cache.Load(r.db, "addTeam", addTeamQuery)
	if err != nil {
		return err
	}

	if tx != nil {
		stmt = tx.NamedStmt(stmt)
	}

	arg := struct {
		LeagueId string `db:"leagueId"`
		TeamId   string `db:"teamId"`
	}{
		LeagueId: leagueId,
		TeamId:   teamId,
	}

	res, err := stmt.Exec(arg)
	if err != nil {
		return err
	}

	if n, err := res.RowsAffected(); err != nil {
		return err
	} else if n != 1 {
		return fmt.Errorf("Expected 1 affected row, found %d", n)
	}

	return nil
}

func (r *repo) clearTeamsFromLeague(leagueId string, tx *sqlx.Tx) error {
	stmt, err := cache.Load(r.db, "clearTeams", clearTeamsQuery)
	if err != nil {
		return err
	}

	if tx == nil {
		return fmt.Errorf("tx is required")
	}

	stmt = tx.NamedStmt(stmt)
	arg := struct{ Id string }{Id: leagueId}
	_, err = stmt.Exec(arg)
	return err
}

func (r *repo) dtosToLeagues(dtos []leagueTeamDto) []League {
	leagues := make(map[string]*League)
	for _, d := range dtos {
		if l, ok := leagues[d.PublicId]; ok {
			if d.TeamId.Valid {
				l.Teams = append(l.Teams, d.TeamId.String)
			}
		} else {
			new := League{
				PublicId:    d.PublicId,
				Name:        d.Name,
				Description: d.Description,
				Location:    d.Location,
				Division:    d.Division,
				Gender:      d.Gender,
				StartDate:   d.StartDate,
				EndDate:     d.EndDate,
				Teams:       []string{},
			}

			if d.TeamId.Valid {
				new.Teams = []string{d.TeamId.String}
			}

			leagues[d.PublicId] = &new
		}
	}

	res := make([]League, 0, len(leagues))
	for _, v := range leagues {
		res = append(res, *v)
	}

	return res
}

func idStruct(id int) interface{} {
	return struct {
		Id int `db:"id"`
	}{
		Id: id,
	}
}
