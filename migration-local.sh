#!/bin/bash
export MIGRATION_DIR=./migrations
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="note-service"
export DB_USER="note-service-user"
export DB_PASSWORD="note-serice-password"
export DB_SSL=disable
export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" goose up -v