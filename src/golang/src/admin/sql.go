package admin

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type account struct {
	identifier   string
	passwordHash string
	connection   string
}

func (r *Router) findUserByIdentifier(identifier string) (*account, error) {
	db, err := sqlx.Open("mysql", r.ConnectionString)
	if err != nil {
		fmt.Printf("Failed to connect to admin db: %s", err)
		return nil, err
	}

	account := account{
		identifier: identifier,
	}

	row := db.QueryRow(findLoginsQuery, identifier)
	if err = row.Scan(&account.passwordHash, &account.connection); err != nil {
		fmt.Printf("Failed to query row from admin db: %s", err)
		return nil, err
	}

	return &account, nil
}

var findLoginsQuery string = `
SELECT pl.passwordHash
     , HEX(c.publicId) as connection
  from principalLogin pl
  join principal p on p.id = pl.principalId
  join dbConnection c on c.id = p.dbConnectionId
 WHERE pl.identifier = ?
`
