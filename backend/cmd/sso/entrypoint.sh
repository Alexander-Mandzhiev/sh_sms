#!/bin/sh
set -e

echo "DB_HOST: client_database"
echo "DB_USER: ${SSO_DB_USER}"
echo "DB_NAME: ${SSO_DB_NAME}"

CONNECTION_STRING="postgres://${SSO_DB_USER}:${SSO_DB_PASSWORD}@sso_database:5432/${SSO_DB_NAME}?sslmode=disable"

wait_for_db() {
  until psql "$CONNECTION_STRING" -c "SELECT 1" >/dev/null 2>&1; do
    echo "Waiting for sso database..."
    sleep 2
  done
}

wait_for_db

echo "Applying sso migrations..."
goose -dir migrations/sso postgres "$CONNECTION_STRING" up

exec ./main