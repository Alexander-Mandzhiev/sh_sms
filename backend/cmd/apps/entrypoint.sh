#!/bin/sh
set -e


echo "DB_HOST: apps_database"
echo "DB_USER: ${APPS_DB_USER}"
echo "DB_NAME: ${APPS_DB_NAME}"

CONNECTION_STRING="postgres://${APPS_DB_USER}:${APPS_DB_PASSWORD}@apps_database:5432/${APPS_DB_NAME}?sslmode=disable"

wait_for_db() {
  until psql "$CONNECTION_STRING" -c "SELECT 1" >/dev/null 2>&1; do
    echo "Waiting for database..."
    sleep 2
  done
}

wait_for_db

echo "Applying migrations..."
goose -dir migrations/apps postgres "$CONNECTION_STRING" up

exec ./main