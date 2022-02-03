migrate-up:
	migrate -path ./migrations -database ${DB_URL} up