# Commish
##### An app for managing player statistics and scheduling leagues and tournaments of various sports and competitions.

## Api Design
### Connections
---
The app needs to support multiple database connections.  This will allow each client to manage and own their own data, and potentially extend the data model.  Each client's connection is stored with a non-guessable ID in a master database.  Each connection in the master database points to another database that MUST define all of the schema in sql/schema.  To obtain a connection, a client must POST to /admin/logins with the appropriate credentials.  This will return a connection id and a JWT.  The connection id is public data, and can safely be hard-coded into a front-end implementation if necessary.  

All non-admin operations require the client to pass in a connection id.  This should be included in the `X-COMMISH-CONNECTION` header.

### Resources
#### Players
---
The player resource defines a player.  
Endpoints:

* GET /players - list of all players
* POST /players - Create a new player
* GET /players/{id} - Fetch a detailed view of a specific player
* PUT /players/{id} - Replace a player
* PATCH /players/{id} - Update one or more specific fields of a player
* DELETE /players/{id} - Remove a player

Subcollections:

* GET /players/{id}/teams
* GET /players/{id}/leagues
* GET /players/{id}/tournaments

#### Teams
---
A team is a collection of players.  It may be entered into leagues and tournaments
Endpoints:

* GET /teams - list of all teams
* POST /teams - create a new team
* GET /teams/{id} - detailed view of a team
* PUT /teams/{id} - replace a team
* PATCH /players/{id} - Update one or more specified fields
* DELETE /teams/{id} - delete a team

Subcollections:

* GET /teams/{id}/players
* POST /teams/{id}/players - add a player to this team
* GET /teams/{id}/leagues
* GET /teams/{id}/tournaments

#### Matches
---
A match represents a logical contest between two teams.  It should contain one or more games.

* GET /matches/{id}
* PATCH /matches/{id}
* DELETE /matches/{id}
* GET /matches/{id}/games
* POST /matches/{id}/games

#### Games
---
A game represents a contest between two teams.  One or more games may be required to resolve a match.

* GET /games/{id}
* PATCH /games/{id}
* GET /games/{id}/teams

#### Leagues
---
A league is a collection of teams, and a collection of games between those teams.  Games may be
scheduled for the future, or may have already been played and have result data.

* GET /leagues - list of all leagues
* POST /leagues - create a new league
* GET /leagues/{id} - detailed view of a league
* PUT /leagues/{id}
* DELETE /leagues/{id}

Subcollections:

* GET /leagues/{id}/teams
* POST /leagues/{id}/teams - add a team to this league
* GET /leagues/{id}/tournaments - get all tournaments associated with this league
* POST /leagues/{id}/tournaments - add a tournament to this league
* GET /leagues/{id}/matches
* POST /leagues/{id}/matches - add a match to this league

#### Tournaments
---
A tournament is a group of teams and games with a clear, single winner.  The winner is determined by the result of the final game.  A tournament may optionally belong to a league (ie end-of-season tournament).  A league could also consist entirely of tournaments, for instance to represent a tournament season.

* PATCH /tournaments/{id}
* DELETE /tournaments/{id}

* GET /tournaments/{id}/matches
* POST /tournaments/{id}/matches
* GET /tournaments/{id}/teams
* POST /tournaments/{id}/teams