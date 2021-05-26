GIT_HASH ?= $(shell git describe --always --dirty --match "v[0-9]*")

PROJECT_NAME ?= 3-tier-counter-backend
HASH_TAG     ?= gcr.io/sada-aaron-brock/$(PROJECT_NAME):$(GIT_HASH)
NAME_TAG     ?= gcr.io/sada-aaron-brock/$(PROJECT_NAME):latest

export DOCKER_BUILDKIT = 1
pull:
	-docker pull $(HASH_TAG)
	-docker pull $(NAME_TAG)

build: pull
	docker build \
		-t $(HASH_TAG) \
		-t $(NAME_TAG) \
		--cache-from $(HASH_TAG) \
		--cache-from $(NAME_TAG) \
	 	--build-arg BUILDKIT_INLINE_CACHE=1 \
		.

push: build
	docker push $(HASH_TAG)
	docker push $(NAME_TAG)

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