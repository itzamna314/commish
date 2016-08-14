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
	 , HEX(t.publicId) as teamPublicId
  FROM league l
  JOIN division d on d.id = l.divisionId
  JOIN genderType g on g.id = l.genderId
  LEFT JOIN leagueTeam lt on lt.leagueId = l.id
  LEFT JOIN team t on t.id = lt.teamId
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
	 , HEX(t.publicId) as teamPublicId
  FROM league l
  JOIN division d on d.id = l.divisionId
  JOIN genderType g on g.id = l.genderId
  LEFT JOIN leagueTeam lt on lt.leagueId = l.id
  LEFT JOIN team t on t.id = lt.teamId
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

	updateQuery = `
UPDATE league l
  LEFT JOIN genderType g on g.name = :gender
  LEFT JOIN division d on d.name = :division
   SET l.name = COALESCE(:name, l.name)
     , l.description = COALESCE(:description, l.description)
	 , l.location = COALESCE(:location, l.location)
	 , l.genderId = COALESCE(g.id, l.genderId)
	 , l.divisionId = COALESCE(d.id, l.divisionId)
	 , l.startDate = COALESCE(NULLIF(:startDate, ''), l.startDate)
	 , l.endDate = COALESCE(NULLIF(:endDate, ''), l.endDate)
     , l.modifiedOn = CURRENT_TIMESTAMP
	 , l.modifiedBy = 'leagues/updateQuery'
 WHERE HEX(l.publicId) = :publicId;
`

	addTeamQuery = `
INSERT INTO leagueTeam (leagueId, teamId, createdOn, createdBy)
     SELECT l.id
	      , t.id
		  , CURRENT_TIMESTAMP
		  , 'leagues/addTeamQuery'
       FROM league l
	   JOIN team t
	  WHERE HEX(l.publicId) = :leagueId
	    AND HEX(t.publicId) = :teamId
`

	clearTeamsQuery = `
DELETE lt
  FROM leagueTeam lt
  JOIN league l on l.id = lt.leagueId
 WHERE HEX(l.publicId) = :id
`
)
