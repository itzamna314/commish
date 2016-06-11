#!/bin/sh

if [ $# -ne 2 ]; then
    echo "usage: $0 <homeTeam> <awayTeam>"
    exit 1
fi

HOME=$1
AWAY=$2

curl -X POST $HOST/api/matches -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"awayTeam":"'$AWAY'", "homeTeam":"'$HOME'"}'
