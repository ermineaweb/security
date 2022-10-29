#!/bin/bash

# zap in browser
# https://www.zaproxy.org/docs/docker/webswing/
docker run --rm -it \
-p 8080:8080 \
-p 8090:8090 \
-u zap \
owasp/zap2docker-weekly \
zap-webswing.sh
