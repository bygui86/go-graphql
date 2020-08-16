
# VARIABLES
# -


# GLOBAL
export GO111MODULE=on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-postgres :		## Run PostgreSQL in a container
	docker run -d --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres

## application

build :		## Build application
	go build

run :		## Run application from source code
	godotenv -f local.env go run main.go

install-gqlgen :		## Install gqlgen
	#go get github.com/99designs/gqlgen
	go get github.com/99designs/gqlgen@v0.11.3

init-graphql :		## Initialise gqlgen for GraphQL
	go run github.com/99designs/gqlgen init

generate-graphql :		## Generate gqlgen code for GraphQL based on gqlgen.yml
	rm -f graph/schema.resolvers.go
	go run github.com/99designs/gqlgen generate

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
