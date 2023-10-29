MYSQL_HOST=mysql
MYSQL_USER=root
MYSQL_PASS=password
MYSQL_NAME=tweets
MYSQL_PORT=3306

DSN := 'mysql://$(MYSQL_USER):$(MYSQL_PASS)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_NAME)'
DOCKER_COMPOSE_EXEC_GO := docker-compose exec go
CREATE_MIGRATE = $(DOCKER_COMPOSE_EXEC_GO) migrate create -ext sql -dir db/migrations -seq $(MIGRATION_NAME)
MIGRATE = $(DOCKER_COMPOSE_EXEC_GO) migrate -path db/migrations -database $(DSN)

# launch or shutdown container
up:
	docker-compose up
down:
	docker-compose down

# enter the container
go_exec:
	docker-compose exec go /bin/sh

db_exec:
	docker-compose exec mysql mysql -uroot -ppassword

# migrate DB
add-migration:  # please pass 'MIGRATION_NAME=${arg}'
	$(CREATE_MIGRATE)
migrateup:
	$(MIGRATE) -verbose up
migrateup1:
	$(MIGRATE) -verbose up 1
migratedown:
	$(MIGRATE) -verbose down
migratedown1:
	$(MIGRATE) -verbose down 1
migrateforce:	# please pass 'MIGRATION_VERSION=${arg}'
	$(MIGRATE) force $(MIGRATION_VERSION)

# upload movie file to vimeo in local environment
upload: # please pass 'MOVIE_FILE_PATH=${arg}'
	$(DOCKER_COMPOSE_EXEC_GO) go run /app/cmd/musemo_batch/upload_video.go -filename=$(MOVIE_FILE_PATH)

# input initial data
seeds:
	docker-compose exec -T mysql mysql -uroot -proot musemo < db/seeds/seeds.sql

# input test user data
test_users:
	docker-compose exec -T mysql mysql -uroot -proot musemo < db/test_users/users.sql
