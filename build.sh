#! /bin/bash

set -e
set -x

rm txterv8 || true
docker rm $(docker ps -aq) || true
docker rmi txterv8:latest || true

go build -o txterv8 main.go
docker build . -f Dockerfile -t txterv8:latest
