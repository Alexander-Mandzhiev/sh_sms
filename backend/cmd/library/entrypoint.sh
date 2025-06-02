#!/bin/sh
set -e

echo "DB_HOST: library_database"
echo "DB_USER: ${LIB_DB_USER}"
echo "DB_NAME: ${LIB_DB_NAME}"

CONNECTION_STRING="postgres://${LIB_DB_USER}:${LIB_DB_PASSWORD}@library_database:5432/${LIB_DB_NAME}?sslmode=disable"

wait_for_db() {
  until psql "$CONNECTION_STRING" -c "SELECT 1" >/dev/null 2>&1; do
    echo "Waiting for library database..."
    sleep 2
  done
}

wait_for_db

echo "Applying library migrations..."
goose -dir migrations/library postgres "$CONNECTION_STRING" up

exec ./main