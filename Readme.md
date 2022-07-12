colima start
docker run --name harmony-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=harmony-db postgres
sh database/run_migrations.sh
go run main.go