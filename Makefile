.PHONY: all

service := golangservicel0

start-servers:
	@echo "  >  Building binary..."
	sudo docker-compose up -d postgresql

init-postgresql: start-servers
	sudo docker-compose exec -T postgresql psql -U default < pkg/repository/postgres/postgresql.sql

connect-to-postgresql: start-servers
	sudo docker-compose exec postgresql psql -U default

build:
	sudo docker-compose build --no -cache ${service}

start: start-servers
	sudo docker-compose --compatibility up --build -d ${service}

run:
	sudo docker-compose --compatibility up --build ${service}

clean:
	sudo docker-compose down --volumes --remove-orphans

