#!/bin/bash
if [ $# -ne 1 ]; then
    echo "usage: $0 <publicId>"
    exit 1
fi

PUBLIC_ID=$1

curl -X GET -H "$COMMISH_CONN" $HOST/api/leagues/$PUBLIC_ID
