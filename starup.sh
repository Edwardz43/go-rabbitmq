#!/bin/bash
#This bash script will create a docker container running rabbitmq. 
#Accessing the management website at http://localhost:45726 with guest/guest(user/passwd).
docker run -d --name myrabbit -e RABBITMQ_DEFAULT_USER=docker1 -e RABBITMQ_DEFAULT_PASS=P@ssw0rd -p 45672:15672 -p 35672:5672 rabbitmq:3-management
