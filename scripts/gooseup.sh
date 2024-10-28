cd server/sql/schema
goose postgres $DB_CONNECTION_STRING up
cd ../../..