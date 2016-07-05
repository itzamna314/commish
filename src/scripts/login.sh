# curl -X POST $HOST/api/admin/logins -H "Content-Type: application/json" -d '{"identifier":"ava","password":"ava"}'

# curl -X POST $HOST/api/admin/logins -H "Content-Type: application/json" -d '{"identifier":"commish","password":"commish"}'

curl -X POST $HOST/api/admin/logins -H "Content-Type: application/json" -d '{"identifier":"commishLocal", "password":"commish"}'


