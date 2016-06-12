package matches

var (
	listQuery = `
SELECT HEX(m.publicId) as publicId
     , HEX(ht.publicId) as homeTeamId
	 , HEX(at.publicId) as awayTeamId
	 , gst.Name as state
  FROM ` + "`match`" + ` m
  JOIN gameStateType gst on gst.id = m.stateId
  LEFT JOIN team ht on ht.id = m.homeTeamId
  LEFT JOIN team at on at.id = m.awayTeamId
	`

	fetchQuery = `
SELECT HEX(m.publicId) as publicId
     , m.homeTeamId
	 , m.awayTeamId
	 , gst.Name as state
  FROM ` + "`match`" + ` m
  JOIN gameStateType gst on gst.id = m.stateId
 WHERE HEX(m.publicId)=:id
`

	fetchPrivateQuery = `
SELECT HEX(m.publicId) as publicId
     , HEX(ht.publicId) as homeTeamId
	 , HEX(at.publicId) as awayTeamId
	 , gst.name as state 
  FROM ` + "`match`" + ` m
  LEFT JOIN team ht on ht.id = m.homeTeamId
  LEFT JOIN team at on at.id = m.awayTeamId
  JOIN gameStateType gst on gst.id = m.stateId
 WHERE m.id=:id
`

	createQuery = `
INSERT INTO ` + "`match`" + ` (awayTeamId, homeTeamId, stateId, createdOn, createdBy)
	SELECT at.id 
	     , ht.id 
		 , gst.id
		 , CURRENT_TIMESTAMP
		 , 'matches/createQuery'
	  FROM gameStateType gst
	  LEFT JOIN team at on HEX(at.publicId)=:awayTeamId
	  LEFT JOIN team ht on HEX(ht.publicId)=:homeTeamId
	 WHERE gst.name = :state;
`

	replaceQuery = `
UPDATE ` + "`match`" + ` m
  JOIN gameStateType gst on gst.name = :state
  JOIN team ht on HEX(ht.publicId) = :homeTeamId
  JOIN team at on HEX(at.publicId) = :awayTeamId
   SET m.awayTeamId=at.id
     , m.homeTeamId=ht.id
	 , m.stateId=gst.id
	 , m.modifiedOn = CURRENT_TIMESTAMP
	 , m.modifiedBy = 'matches/replaceQuery'
 WHERE HEX(m.publicId) = :publicId;
`
)
