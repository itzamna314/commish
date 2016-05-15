package players

const (
	listPlayersQuery = `
SELECT publicId
     , name
	 , age
	 , gender
  FROM players`
)
