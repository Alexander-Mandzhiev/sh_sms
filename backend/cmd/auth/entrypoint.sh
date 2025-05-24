#!/bin/sh
set -e

echo "DB_HOST: auth_database"
echo "DB_USER: ${AUTH_DB_USER}"
echo "DB_NAME: ${AUTH_DB_NAME}"

CONNECTION_STRING="postgres://${AUTH_DB_USER}:${AUTH_DB_PASSWORD}@auth_database:5432/${AUTH_DB_NAME}?sslmode=disable"

wait_for_db() {
  until psql "$CONNECTION_STRING" -c "SELECT 1" >/dev/null 2>&1; do
    echo "Waiting for auth database..."
    sleep 2
  done
}

wait_for_db

echo "Applying auth migrations..."
goose -dir migrations/auth postgres "$CONNECTION_STRING" up

exec ./main