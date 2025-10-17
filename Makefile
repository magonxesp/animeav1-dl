.PHONY: build-docker-image test

build-docker-image:
	export VERSION="$$(date +%Y-%m-%d.%H%M%S)"; \
	docker build -t "magonx/animeav1-dl:$$VERSION" . --platform linux/amd64,linux/arm64; \
	echo "magonx/animeav1-dl:$$VERSION"

test:
	go test -v ./internal/*
