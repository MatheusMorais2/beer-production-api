migrate_cmd = migrate -source file://adapters/db/migrations -database "postgresql://beer-production-api:password@localhost:5432/beer-production-api-database?sslmode=disable" -verbose
create_migration_cmd = migrate create -ext sql -dir adapters/db/migrations -seq

migration_up:
	$(migrate_cmd) up

migration_down:	
	$(migrate_cmd) down

migration_fix:	
	$(migrate_cmd) force ${DB_MIGRATION_OK}	

create_migration:
	$(create_migration_cmd) ${name}
