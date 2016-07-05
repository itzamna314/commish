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

	findFlagPlayer = 1

	findQuery = `
SELECT HEX(t.publicId) as publicId
     , t.name
  FROM team t
 WHERE (
	   $FLAGS$ & 1 = 0 
	   OR EXISTS (SELECT 1 
	                FROM playerTeam pt
					JOIN player p on p.id = pt.playerId
				   WHERE HEX(p.publicId) = :playerId
				     AND pt.teamId = t.id)
	   )
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
