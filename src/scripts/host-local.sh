export HOST=http://localhost:8080
RESULT=`curl -X POST $HOST/api/admin/logins -H "Content-Type: application/json" -d '{"identifier":"commishLocal","password":"commish"}'`

export COMMISH_CONN='X-COMMISH-CONNECTION: '`echo $RESULT | jq -r '.user.connection'`
export COMMISH_AUTH='Authorization: Bearer '`echo $RESULT | jq -r '.user.token'`
