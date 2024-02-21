#!/bin/bash

URL=http://${HOST}:${PORT}/health
echo "$URL:"
curl -s $VERBOSE http://${HOST}:${PORT}/health | jq
