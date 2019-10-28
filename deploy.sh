#!/bin/sh

docker-compose stop
docker-compose rm
docker-compose up -d --build

echo "done docker compose, start yarn...."

cd frontend && yarn && yarn serve

