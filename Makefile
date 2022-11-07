build: clean
	docker-compose build

up:
	docker compose up

down:
	docker compose down --remove-orphans --volumes

clean:
	docker compose down --remove-orphans
	docker rm --force app
	docker rm --force mysql-db
	docker rmi -f go-bookstore
	docker rmi -f mysql:8.0.22