#! /bin/bash

set -e
set -x

./build.sh

docker run -p 8888:8888 txterv8:latest
