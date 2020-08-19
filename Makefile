
# VARIABLES
# -


# GLOBAL
export GO111MODULE=on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

run-mysql :		## Start MySQL container
	docker run -d --name mysql \
		-e MYSQL_ROOT_PASSWORD=supersecret \
		-p 3306:3306 \
		mysql

run-postgres :		## Start PostgreSQL container
	docker run -d --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres

## applications

run-server :		## Start GraphQL server
	cd server/ && \
		make run

run-client :		## Start GraphQL client
	@echo "NOT IMPLEMENTED"
	#cd client/ && \
#		make run

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
