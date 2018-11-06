#!/bin/bash
set -e
# move to root directory
cd ..
# docker build
docker build -t dongfeng-websocket . -f DockerFile
echo "Bye"