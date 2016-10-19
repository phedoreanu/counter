#!/bin/bash

if [ ! -d "$DIRECTORY" ]; then
  go get -u github.com/tsenart/vegeta
fi

echo "GET http://${HOSTNAME}/get" | vegeta attack -duration=5s | tee results.bin | vegeta report
cat results.bin | vegeta report -reporter=plot > plot.html
