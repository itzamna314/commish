package players

const (
	listQuery = `
SELECT HEX(p.publicId) as playerPublicId
     , p.name as playerName
	 , p.age as playerAge
	 , g.name as playerGender
	 , HEX(t.publicId) as teamPublicId
  FROM player p
  JOIN genderType g on g.id = p.genderId
  LEFT JOIN playerTeam pt on pt.playerId = p.id
  LEFT JOIN team t on t.id = pt.teamId
`

	fetchQuery = `
SELECT HEX(p.publicId) as publicId
     , p.name
	 , p.age
	 , g.name as gender
  FROM player p
  JOIN genderType g on g.id = p.genderId
 WHERE HEX(p.publicId)=:id
`

	fetchPrivateQuery = `
SELECT HEX(p.publicId) as publicId
	 , p.name
     , p.age
     , g.name as gender
  FROM player p
  JOIN genderType g on g.id = p.genderId
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
)
