#!/bin/bash

APP="websec"
LOCAL_PATH_RESULTS=$(pwd)/results
TARGET="http://localhost:3000"

mkdir -p "$LOCAL_PATH_RESULTS"

# docker build -t "$APP":latest .

# docker run --rm -it \
# --net=host \
# -v "$LOCAL_PATH_RESULTS":/security/results \
# "$APP":latest $TARGET

# zap basic scan
# https://www.zaproxy.org/docs/docker/baseline-scan/
# docker run --rm -it \
# --net=host \
# -v "$LOCAL_PATH_RESULTS":/zap/wrk:rw \
# owasp/zap2docker-weekly \
# zap-baseline.py -t "$TARGET" -w zap_report_basic.md

# zap full scan
# https://www.zaproxy.org/docs/docker/full-scan/
docker run --rm -it \
--net=host \
-v "$LOCAL_PATH_RESULTS":/zap/wrk:rw \
owasp/zap2docker-weekly \
zap-full-scan.py -t "$TARGET" -w zap_report_full.md

# zap API scan
# https://www.zaproxy.org/docs/docker/api-scan/
# docker run --rm -it \
# --net=host \
# -v "$LOCAL_PATH_RESULTS":/zap/wrk:rw \
# owasp/zap2docker-weekly \
# zap-api-scan.py -t "$TARGET" -f openapi -a -w zap_report_api.md