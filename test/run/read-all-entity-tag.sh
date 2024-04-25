#!/bin/bash

HOST=localhost
PORT=8080
URL=http://${HOST}:${PORT}/api/entity/tags
CMD="curl $URL"
echo Running commad: $CMD
eval $CMD
