curl -X POST $HOST/api/leagues -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Mens summer 2s", "location":"sandbox", "description":"for noobs!", "division": "B", "gender":"male","startDate":"2016-04-20", "endDate":"2016-06-01"}'