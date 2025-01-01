.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest

push: ## push docker image to artifact registry
	docker buildx build --platform linux/amd64 \
  --target deploy \
  -t asia-southeast1-docker.pkg.dev/plasma-renderer-446307-u5/task-api-repositry/gotodo:latest \
  --push .

auth-repo: ## authorize artifact registry
	gcloud auth configure-docker asia-southeast1-docker.pkg.dev

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'