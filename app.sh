#!/bin/bash

set -e

case $1 in
  run)
    export DATABASE_URL=postgres://aether_api:not_the_password@localhost:5432/aether
    go run cmd/web-server/main.go
    ;;
  *)
    echo "Usage ./app.sh {run}"
  ;;
esac
