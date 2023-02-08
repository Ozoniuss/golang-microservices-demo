#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER ports_user PASSWORD 'userpassword';
	CREATE DATABASE ports_db;
	GRANT ALL ON DATABASE ports_db TO ports_user;
EOSQL

psql_command() {
    psql --username ports_user --dbname ports_db -f "$1"
}

for f in /scripts/migrate/*.up.sql
do
    psql_command ${f}
done