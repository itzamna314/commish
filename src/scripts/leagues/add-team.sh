if [[ $# -ne 2 ]]; then
    echo "usage: $0 <leagueId> <teamId>"
    exit 1
fi

curl -X PATCH $HOST/api/leagues/$1 -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name": "Updated teams", "teams":["'$2'"]}'
