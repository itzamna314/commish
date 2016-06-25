package divisions

const (
	createQuery = `
INSERT INTO division (name, createdOn, createdBy) VALUES
	(:name, CURRENT_TIMESTAMP, 'divisions/create')
`

	fetchQuery = `
SELECT HEX(publicId) as publicId
    , name
 FROM division
WHERE name=:name;
`
)
