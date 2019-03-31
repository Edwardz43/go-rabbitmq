#!/bin/bash
## Usage: ./setup.sh <number of producer> <server ip>
## Server IP is usually the Docker gateway IP address, which is 172.17.0.1 by default

REPLICAS=$1
IP=$2
go build --tags "static netgo" -o send send.go
for (( c=0; c<${REPLICAS}; c++ ))
do
    docker run -l rabbitmq -v $(pwd)/send:/send alpine /send -ip=${IP}
done
