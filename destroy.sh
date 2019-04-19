#!/bin/bash
#This bash script will remove the rabbitmq docker container.
docker rm --force $(docker ps -aq --filter "name=myrabbit")
