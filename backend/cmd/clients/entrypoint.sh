#!/bin/sh
set -e

echo "DB_HOST: client_database"
echo "DB_USER: ${CLIENTS_DB_USER}"
echo "DB_NAME: ${CLIENTS_DB_NAME}"

CONNECTION_STRING="postgres://${CLIENTS_DB_USER}:${CLIENTS_DB_PASSWORD}@client_database:5432/${CLIENTS_DB_NAME}?sslmode=disable"

wait_for_db() {
  until psql "$CONNECTION_STRING" -c "SELECT 1" >/dev/null 2>&1; do
    echo "Waiting for clients database..."
    sleep 2
  done
}

wait_for_db

echo "Applying clients migrations..."
goose -dir migrations/clients postgres "$CONNECTION_STRING" up

exec ./main