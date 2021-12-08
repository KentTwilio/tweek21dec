#! /bin/bash

set -e
set -x

./build.sh

docker run -p 80:80 txterv8:latest
