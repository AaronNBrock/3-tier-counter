DOCKER_TAG   ?= $(shell git describe --always --dirty --match "v[0-9]*")
PROJECT_NAME ?= 3-tier-counter-backend
TAG          ?= gcr.io/sada-aaron-brock/$(PROJECT_NAME):$(DOCKER_TAG)

export DOCKER_BUILDKIT = 1
pull:
	-docker pull $(TAG)

build: pull
	docker build \
		-t $(TAG) \
		--cache-from $(TAG) \
	 	--build-arg BUILDKIT_INLINE_CACHE=1 \
		.

push: build
	docker push $(TAG)

.PHONY: deploy
deploy:
	kubectl apply --force -k deploy/

.PHONY: deploy
undeploy:
	kubectl delete -k deploy/


#
# Local Development
#
start:
	-@docker network create 3-tier-counter
	docker run --rm -d -p 3000:3000 -e DB_HOST=3-tier-counter-database --network 3-tier-counter --name $(NAME) $(TAG)

logs:
	docker logs --follow $(NAME)

stop:
	docker stop $(NAME)
	@-docker network rm 3-tier-counter

fmt:
	docker run --rm -it --volume $(shell pwd):/app --workdir /app golang:1.15-alpine go fmt