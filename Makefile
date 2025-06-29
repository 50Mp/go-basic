dev:
	go run main.go ./env/dev/.env.auth

prod:
	go run main.go ./env/prod/.env.prod

test:
	go run main.go ./test/prod/.env.test

up:
	podman-compose -f docker-compose.db.yml up --detach
down:
	podman-compose -f docker-compose.db.yml down

start:
	podman-compose -f docker-compose.db.yml start
stop:
	podman-compose -f docker-compose.db.yml stop
ps:
	podman-compose -f docker-compose.db.yml ps
images:
	podman-compose -f docker-compose.db.yml images
	


tidy:
	go mod tidy

.PHONY: dev prod test up down start stop ps images tydi