#!/bin/bash

DATA_FILE="./default_major_data.sql"
USER="admin"
DB_ADDR="localhost"
DB_PORT="6432"
DB="major"

DIR=$(dirname "$0")

psql -U $USER -d $DB -h $DB_ADDR -p $DB_PORT -f $DIR/$DATA_FILE
