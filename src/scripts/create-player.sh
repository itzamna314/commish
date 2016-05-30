curl -X POST $HOST/api/players -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Gerald", "age":42, "gender":"male"}'
