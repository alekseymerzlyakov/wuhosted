#!/bin/bash
login=$1
key=$2
env=$3
endpoint=$4
enc=`echo "$login:$key" | base64`
auth=`curl -s -X GET https://portal.$env.wuamerigo.com/admin/api/1/oauth2/token -H "Authorization: Basic $enc" -H 'Content-Type: application/json'`
echo "$auth"
token=`echo $auth | jq -r '.["Access-Token"]'`
echo $token
rez=`curl -s -X GET https://portal.$env.wuamerigo.com/admin/api/1/$4 -H "Authorization: Bearer $token" -H 'Content-Type: application/json'`
echo "$rez"
