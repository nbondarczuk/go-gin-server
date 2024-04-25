#!/bin/bash

HOST=localhost
PORT=8080
URL=http://${HOST}:${PORT}/api/entity/tags
CMD="curl $URL"
echo -n Running command: $CMD " - result: "
eval $CMD
