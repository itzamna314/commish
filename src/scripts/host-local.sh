export HOST=http://localhost:8080
RESULT=`curl -X POST $HOST/admin/logins -H "Content-Type: application/json" -d '{"identifier":"ava","password":"ava"}'`

export COMMISH_CONN='X-COMMISH-CONNECTION: '`echo $RESULT | jq -r '.user.connection'`
export COMMISH_AUTH='Authorization: Bearer '`echo $RESULT | jq -r '.user.token'`
