#!/bin/sh

docker-compose stop
docker-compose rm
docker-compose up -d --build

echo "done"

