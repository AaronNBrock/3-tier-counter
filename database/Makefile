TAG  ?= postgres:12.4-alpine
NAME := 3-tier-counter-database

start:
	-@docker network create 3-tier-counter
	docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD=password --network 3-tier-counter --name $(NAME) $(TAG)

logs:
	docker logs --follow $(NAME)

stop:
	docker stop $(NAME)
	@-docker network rm 3-tier-counter

# To avoid errors
build:
push:
fmt:
