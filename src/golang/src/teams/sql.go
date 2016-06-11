package teams

const (
	listQuery = `
SELECT HEX(t.publicId) as publicId
     , t.name
  FROM team t
`

	fetchQuery = `
SELECT HEX(t.publicId) as publicId
     , t.name
  FROM team t 
 WHERE HEX(t.publicId)=:id
`

	fetchPrivateQuery = `
SELECT HEX(t.publicId) as publicId
	 , t.name
  FROM team t
 WHERE t.id=:id
`

	createQuery = `
INSERT INTO team(Name, CreatedOn, CreatedBy)
VALUES(:name, CURRENT_TIMESTAMP,'teams/createQuery')
`

	replaceQuery = `
UPDATE team t
   SET t.name = :name
     , t.modifiedOn = CURRENT_TIMESTAMP
	 , t.modifiedBy = 'teams/replaceQuery'
 WHERE HEX(t.publicId) = :publicId;
`
)
