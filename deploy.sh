#! /bin/bash

# 命令行参数
endpoint=$1
config=$2

if [ "$config" == "" ]; then
    config="./conf/dev/"
fi

if [ "$endpoint" == "" ]; then
    endpoint="dashboard"
    echo "********************build project********************"
    cd /application/Gateway
    env GOOS=linux GOARCH=amd64 go build -o dashbord main.go

    echo "********************delete old version********************"
    pid=$(netstat -ntlp | grep ./dashbord | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }' | tail -n 1)
        if [ -n "$pid" ] && [ $pid -gt 0 ]; then
            kill $pid
        fi

    echo "********************restart project********************"
    nohup ./dashbord -endpoint=$endpoint -config=$config > ./dashbord.nohup.log 2>&1 &
fi


if [ "$endpoint" == "server" ]; then
    echo "********************build project********************"
    cd /application/Gateway
    env GOOS=linux GOARCH=amd64 go build -o server main.go

    echo "********************delete old version********************"
    pid=$(netstat -ntlp | grep ./server | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }' | tail -n 1)
        if [ -n "$pid" ] && [ $pid -gt 0 ]; then
            kill $pid
        fi

    echo "********************restart project********************"
    nohup ./server -endpoint=$endpoint -config=$config > ./server.nohup.log 2>&1 &

    echo "********************restart real service********************"
    cd /application/Gateway/services
    env GOOS=linux GOARCH=amd64 go build -o realServer ./main.go
    ppid=$(netstat -ntlp | grep ./realServer | awk '{print $7}' | awk 'match($0, /[0-9]+/) { print substr($0, RSTART, RLENGTH) }' | tail -n 1)
    if [ -n "$ppid" ] && [ $ppid -gt 0 ]; then
        kill $ppid
    fi
    nohup ./realServer > ./realServer.nohup.log 2>&1 &
    sleep 2
    curl "http://127.0.0.1:8080/ping"
    curl -k --insecure "https://127.0.0.1:4433/ping"
    curl "http://127.0.0.1:2001"
    curl "http://127.0.0.1:2002"
    curl "http://127.0.0.1:2003"
fi
