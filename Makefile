.PHONY: rebuild
rebuild:
	docker-compose -p mwh -f deployments/docker-compose.development.yml build --no-cache

.PHONY: build
build:
	docker-compose -p mwh -f deployments/docker-compose.development.yml build

.PHONY: start
start:
	$(MAKE) build
	docker-compose -p mwh -f deployments/docker-compose.development.yml up -d
	$(MAKE) logs

.PHONY: restart
restart:
	$(MAKE) stop
	$(MAKE) start

.PHONY: logs
logs:
	docker-compose -p mwh -f deployments/docker-compose.development.yml logs --tail 50 -f

.PHONY: stop
stop:
	docker-compose -p mwh -f deployments/docker-compose.development.yml down

.PHONY: test
test:
	$(MAKE) stop
	$(MAKE) build
	docker-compose -p mwh -f deployments/docker-compose.development.yml up -d
	curl -i "http://localhost:8000/ping"
	curl -i --user "admin:admin" "http://localhost:8000/films/5"
	docker-compose -p mwh -f deployments/docker-compose.development.yml logs
	$(MAKE) stop
