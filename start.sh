#!/bin/bash
# start.sh using with docker exec
# see deploy/readme.md,  docker exec ginapp-demo ./start.sh
nohup bin/gin-api-common </dev/null >>/logs/ginapp-runtime.log 2>&1 &
