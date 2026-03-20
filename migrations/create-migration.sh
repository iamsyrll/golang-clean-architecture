#!/bin/bash

set -e

# load .env
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

CMD=$1
NAME=$2

if [ "$CMD" = "create" ]; then
  if [ -z "$NAME" ]; then
    echo "Usage: ./migrate.sh create <name>"
    exit 1
  fi

  migrate create -ext sql -dir ./migrations -seq "$NAME"
  exit 0
fi

# default (up, down, version, dll)
migrate -database "$DATABASE_URL" -path ./migrations "$@"