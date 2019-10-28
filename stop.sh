#!/bin/sh

docker-compose stop
docker-compose rm
docker system prune

docker rmi profile_profile
docker rmi profile_postgres

echo "done"

