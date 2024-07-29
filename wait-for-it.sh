#!/bin/bash

# wait-for-it.sh

# Usage: wait-for-it.sh <host>:<port> [<timeout>] [-- command args...]

set -e

TIMEOUT=15
HOST="$1"
shift
COMMAND="$@"

if [[ $HOST == *:* ]]; then
  IFS=":" read -r HOST PORT <<< "$HOST"
else
  echo "Usage: $0 host:port [command args...]"
  exit 1
fi

echo "Waiting for $HOST:$PORT..."

for i in $(seq $TIMEOUT); do
  nc -z $HOST $PORT && break
  echo "Trying again in 1 second..."
  sleep 1
done

if ! nc -z $HOST $PORT; then
  echo "Timeout after $TIMEOUT seconds"
  exit 1
fi

echo "$HOST:$PORT is available"

exec $COMMAND
