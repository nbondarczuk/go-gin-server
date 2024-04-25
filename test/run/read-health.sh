#!/bin/bash

PORT=8080
HOST=localhost
URL=http://${HOST}:${PORT}/system/health
CMD="curl $URL"
echo -n Running command: $CMD " - result: "
eval $CMD
