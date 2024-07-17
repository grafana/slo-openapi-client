.PHONY: generate-go-client
generate-go-client: 
	./scripts/generate.sh

.PHONY: generate-clients
generate-clients: generate-go-client

.PHONY: golangci-lint
golangci-lint:
	docker run \
		--rm \
		--volume "$(shell pwd):/src" \
		--workdir "/src/go" \
		golangci/golangci-lint:v1.55 golangci-lint run ./... -v
