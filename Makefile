help:
	@echo 'Management commands for go-bookstore:'
	@echo '    make build           Compile the project.'
	@echo '    make up              Start the service.'
	@echo '    make clean      	    Clean project.'
	@echo '    make test            Run tests on a compiled project.'
	@echo

build:
	docker-compose -f compose/local.yml build

up:
	docker-compose -f compose/local.yml up

down:
	docker-compose down -f compose/local.yml --remove-orphans
