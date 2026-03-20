#!/bin/bash

set -e

# load .env (robust)
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

if [ -z "$DATABASE_URL" ]; then
  echo "DATABASE_URL not set"
  exit 1
fi

migrate -database "$DATABASE_URL" -path ./migrations "$@"