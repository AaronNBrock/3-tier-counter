DOCKER_TAG   ?= $(shell git describe --always --dirty --match "v[0-9]*")
PROJECT_NAME ?= 3-tier-counter-frontend
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

# Local Dev
start:
	-@docker network create 3-tier-counter
	docker run --rm -d -p 80:80 -e API_ENDPOINT=3-tier-counter-backend:3000 --network 3-tier-counter --name $(NAME) $(TAG)

logs:
	docker logs --follow $(NAME)

stop:
	docker stop $(NAME)
	@-docker network rm 3-tier-counter