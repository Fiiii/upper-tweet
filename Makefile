# REQUIRED SECTION
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
# END OF REQUIRED SECTION

.PHONY: help dependencies up start stop restart status ps clean

DOCKER_COMPOSE = docker-compose
TESTS_DIR = tests
DB_CONNECTION = postgres://admin:password@db_postgres:5432/tweet-db?sslmode=disable
TEST_DB_CONNECTION = postgres://admin:postgres@tweet-db:5432/tweet-db-test?sslmode=disable
RUN_MIGRATION = run --rm migrate -source file://migrations -database

DB = tweet-db
DB_CONTAINER_HOST = db_postgres
DB_CONTAINER_PORT = 5432
DB_CONTAINER_USER = admin
DB_PASSWORD=password

build: ## Build containters
	docker-compose -f docker-compose.yml up --build

up:
	$(DOCKER_COMPOSE) up -d --remove-orphans

logs:
	$(DOCKER_COMPOSE) logs -f

down:
	$(DOCKER_COMPOSE) down

migrate-db-up:
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(DB_CONNECTION) up

migrate-db-down:
	$(DOCKER_COMPOSE) $(RUN_MIGRATION) $(DB_CONNECTION) down 1

psql:
	docker exec -it db_postgres psql $(DB_CONNECTION)