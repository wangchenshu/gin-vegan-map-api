#!/bin/bash

rm -rf gin-vegan-map-api
go build .
docker-compose up -d --build
docker-compose logs --tail=all -f

exit 0

