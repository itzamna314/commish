package leagues

const (
	listQuery = `
SELECT HEX(l.publicId) as publicId
     , l.name
	 , l.description
	 , l.location
	 , d.name as division
	 , g.name as gender
	 , l.startDate
	 , l.endDate
  FROM league l
  JOIN division d on d.id = l.divisionId
  JOIN genderType g on g.id = l.genderId
`

	fetchQuery = `
SELECT HEX(l.publicId) as publicId
     , l.name
	 , l.description
	 , l.location
	 , d.name as division
	 , g.name as gender
	 , l.startDate
	 , l.endDate
  FROM league l
  JOIN division d on d.id = l.divisionId
  JOIN genderType g on g.id = l.genderId
 WHERE HEX(l.publicId)=:id
`

	fetchPrivateQuery = `
SELECT HEX(l.publicId) as publicId
     , l.name
	 , l.description
	 , l.location
	 , d.name as division
	 , g.name as gender
	 , l.startDate
	 , l.endDate
  FROM league l
  JOIN division d on d.id = l.divisionId
  JOIN genderType g on g.id = l.genderId
 WHERE l.id=:id
`

	createQuery = `
INSERT INTO league(name, location, description, divisionId, genderId, startDate, EndDate, createdOn, createdBy)
	SELECT :name
	     , :location
		 , :description
		 , d.id
		 , g.id
		 , :startDate
		 , :endDate
		 , CURRENT_TIMESTAMP
		 , 'leagues/createQuery'
      FROM genderType g
	  JOIN division d on d.name = :division
	 WHERE g.name = :gender
	   AND d.name = :division;
`

	replaceQuery = `
UPDATE league l
  JOIN genderType g on g.name = :gender
  JOIN divisionType d on d.name = :division
   SET l.name = :name
     , l.description = :description
	 , l.location = :location
	 , l.genderId = g.id
	 , l.divisionId = d.id
	 , l.startDate = :startDate
	 , l.endDate = :endDate
     , l.modifiedOn = CURRENT_TIMESTAMP
	 , l.modifiedBy = 'leagues/replaceQuery'
 WHERE HEX(l.publicId) = :publicId;
`
)
