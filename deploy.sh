#! /bin/bash

echo "********************build project********************"
cd /application/Gateway
env GOOS=linux GOARCH=amd64 go build -o gateway main.go

echo "********************delete old version********************"
pid=$(netstat -ntlp | grep ./gateway | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }')
    if [ -n "$pid" ] && [ $pid -gt 0 ]; then
        kill $pid
    fi

echo "********************restart project********************"
nohup ./gateway > ./nohup.log 2>&1 &
