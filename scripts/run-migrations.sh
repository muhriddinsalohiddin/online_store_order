#!/bin/bash
export POSTGRES_USER=muhammad
export POSTGRES_PASSWORD=12345
export POSTGRES_DATABASE=book_order
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path "./migrations"  down
migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path "./migrations"  up
