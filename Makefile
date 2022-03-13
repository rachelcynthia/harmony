db:
	docker rm -f harmony-db || true
	docker run -d --name harmony-db -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=harmony-db postgres

start_db: db
	sh database/run_migrations.sh