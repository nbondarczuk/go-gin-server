#!/bin/bash

PORT=8080
HOST=localhost
URL=http://${HOST}:${PORT}/api/entity/tag
DATA="'{\"label\":\"red\",\"color\": \"red\"}'"
HEADER="\"Content-Type: application/json\""
CMD="curl -H $HEADER -d $DATA $URL"
echo -n Running command: $CMD " - result: "
eval $CMD
