if [[ $# -ne 1 ]]; then
    echo "usage: $0 <publicId>"
    exit 1
fi

curl -X PATCH $HOST/api/leagues/$1 -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Ladies summer 2s","location":"sandbox","description":"for beginners!","division":"BB","gender":"female","startDate":"2016-04-20 00:00:00","endDate":"2016-06-01 00:00:00"}'
