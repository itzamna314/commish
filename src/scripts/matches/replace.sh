if [ $# -ne 1 ]; then
    echo "usage: $0 <publicId>"
    exit 1
fi

curl -v -X PUT $HOST/api/matches/$1 -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"awayTeam":"46E97C24300D11E6A09D409B4CAB7549", "homeTeam":"0FEF5278286C11E6875C81C96BBD740C", "state":"inProgress"}'
