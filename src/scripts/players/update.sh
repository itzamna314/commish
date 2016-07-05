if [ $# -ne 1 ]; then
    echo "usage: $0 <publicId>"
    exit 1
fi

curl -X PATCH $HOST/api/players/$1 -H "Content-Type: application/json" -H "$COMMISH_AUTH" -H "$COMMISH_CONN" -d '{"name":"Felix", "age":"51", "gender":"male", "teams":["81D07C3030DE11E6A09D409B4CAB7549"]}'
