if [ $# -ne 1 ]; then
    echo "usage: $0 <publicId>"
    exit 1
fi

curl -X PUT $HOST/api/players/$1 -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Geraldine", "age":37, "gender":"female"}'
