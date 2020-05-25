# REQUIRED SECTION
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
# END OF REQUIRED SECTION

.PHONY: help dependencies up start stop restart status ps clean

DOCKER_COMPOSE = docker-compose
TESTS_DIR = tests
DB_CONNECTION = postgres://admin:postgres@db-uppergin:5432
TEST_DB_CONNECTION = postgres://admin:postgres@db-uppergin:5432/db-uppergin?sslmode=disable
RUN_MIGRATION = run --rm migrate -source file://migrations -database

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