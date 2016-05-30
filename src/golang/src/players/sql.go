package players

const (
	listQuery = `
SELECT HEX(p.publicId) as publicId
     , p.name
	 , p.age
	 , g.name as gender
  FROM player p
  JOIN genderType g on g.id = p.genderId
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

	replaceQuery = `
UPDATE player p
  JOIN genderType g on g.name = :gender 
   SET p.name = :name
     , p.age = :age
	 , p.genderId = g.id
 WHERE HEX(p.publicId) = :publicId;
`
)
