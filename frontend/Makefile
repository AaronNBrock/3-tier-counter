
TAG  ?= gcr.io/sada-aaron-brock/3-tier-counter-frontend
NAME := 3-tier-counter-frontend

start:
	-@docker network create 3-tier-counter
	docker run --rm -d -p 80:80 -e API_ENDPOINT=3-tier-counter-backend:3000 --network 3-tier-counter --name $(NAME) $(TAG)

logs:
	docker logs --follow $(NAME)

stop:
	docker stop $(NAME)
	@-docker network rm 3-tier-counter

pull:
	-docker pull $(TAG)

build: pull
	docker build --cache-from $(TAG) -t $(TAG) .

push: build
	docker push $(TAG)

fmt:
	docker run --rm -it --volume $(shell pwd):/app --workdir /app golang:1.15-alpine go fmt