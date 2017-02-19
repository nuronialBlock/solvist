#!/bin/bash

echo ">> Waiting for solvist_db to start"
sleep 10
WAIT=0
while ! nc -z solvist_db 27017; do
  sleep 2
  WAIT=$(($WAIT + 1))
  if [ "$WAIT" -gt 25 ]; then
    echo "Error: Timeout wating for solvist_db to start"
    exit 1
  fi
done
echo ">> solvist_db detected"
sh scripts/init.sh