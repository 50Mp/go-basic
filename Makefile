devrun:
	go run main.go ./env/dev/.env.dev

prodrun:
	go run main.go ./env/prod/.env.prod

testrun:
	go run main.go ./test/prod/.env.test

dbup:
	docker compose -f docker-compose.db.yml up -d





# .PHONY: is a special target that tells make that the listed targets are not files
.PHONY: devrun, prodrun, testrun, updb