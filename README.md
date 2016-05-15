# Commish
##### An app for managing player statistics and scheduling leagues and tournaments of various sports and competitions.

## Api Design
---
### Connections
---
The app needs to support multiple database connections.  This will allow each client to manage and own their own data, and potentially extend the data model.  Each client's connection is stored with a non-guessable ID in a master database.  Each connection in the master database points to another database that MUST define all of the schema in sql/schema.

### Resources
---