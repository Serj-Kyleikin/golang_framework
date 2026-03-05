# Root dir where Makefile is located
ROOT := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
COMPOSE := $(ROOT)Docker/docker-compose.yml
COMPOSE_ENV := ROOT=$(ROOT)

.PHONY: up down logs rebuild migrate ps

start:
	cp -n $(ROOT).env.example $(ROOT).env || true
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) up -d --build
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) run --rm migrate

build:
	cp -n $(ROOT).env.example $(ROOT).env || true
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) up --build

up:
	cp -n $(ROOT).env.example $(ROOT).env || true
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) up -d

down:
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) down

ps:
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) ps

migrate:
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) run --rm migrate

logs:
	$(COMPOSE_ENV) docker compose --project-directory $(ROOT) -f $(COMPOSE) logs -f --tail=200