package players

const (
	listQuery = `
SELECT HEX(p.publicId) as publicId
     , p.name
	 , p.age
	 , g.name as gender
	 , HEX(t.publicId) as teamPublicId
  FROM player p
  JOIN genderType g on g.id = p.genderId
  LEFT JOIN playerTeam pt on pt.playerId = p.id
  LEFT JOIN team t on t.id = pt.teamId
`

	fetchQuery = `
SELECT HEX(p.publicId) as playerPublicId
	 , p.name as playerName
     , p.age as playerAge
     , g.name as playerGender
	 , HEX(t.publicId) as teamPublicId
  FROM player p
  JOIN genderType g on g.id = p.genderId
  LEFT JOIN playerTeam pt on pt.playerId = p.id
  LEFT JOIN team t on t.id = pt.teamId
 WHERE HEX(p.publicId)=:id
`

	fetchPrivateQuery = `
SELECT HEX(p.publicId) as playerPublicId
	 , p.name as playerName
     , p.age as playerAge
     , g.name as playerGender
	 , HEX(t.publicId) as teamPublicId
  FROM player p
  JOIN genderType g on g.id = p.genderId
  LEFT JOIN playerTeam pt on pt.playerId = p.id
  LEFT JOIN team t on t.id = pt.teamId
 WHERE p.id=:id
`

	createQuery = `
INSERT INTO player(Name, Age, GenderId, CreatedOn, CreatedBy)
	SELECT :name 
		 , :age
	     , g.id
		 , CURRENT_TIMESTAMP
		 , 'players/createQuery'
	  FROM genderType g
	 WHERE g.name = :gender;
`

	updateQuery = `
UPDATE player p
  JOIN genderType g on g.name = :gender 
   SET p.name = COALESCE(:name, p.name)
     , p.age = COALESCE(:age, p.age)
	 , p.genderId = COALESCE(g.id, p.genderId)
	 , p.modifiedOn = CURRENT_TIMESTAMP
	 , p.modifiedBy = 'players/updateQuery'
 WHERE HEX(p.publicId) = :publicId;
`

	addToTeamQuery = `
INSERT INTO playerTeam (playerId, teamId, createdOn, createdBy) 
	 SELECT p.id
	      , t.id
		  , CURRENT_TIMESTAMP
		  , 'players/addToTeamQuery'
	   FROM player p
	   JOIN team t
	  WHERE HEX(p.publicId) = :playerId
	    AND HEX(t.publicId) = :teamId
`

	clearTeamsQuery = `
DELETE pt 
  FROM playerTeam pt
  JOIN player p on p.id = pt.playerId
 WHERE HEX(p.publicId) = :id
`
)
