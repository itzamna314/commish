PUBLIC_ID=""

if [ ! -t 0 ]; then
    read PUBLIC_ID
elif [ $# -ne 1 ]; then 
    echo "usage: $0 <publicId>"
    exit 1
else
    PUBLIC_ID=$1
fi

curl -X PUT $HOST/api/teams/$PUBLIC_ID -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Unprotected sets"}'
