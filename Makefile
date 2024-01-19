openapi-generator-cli.jar:
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar -O openapi-generator-cli.jar

.PHONY: generate-go-client
generate-go-client: openapi-generator-cli.jar
	GIT_USER_ID=grafana java -jar openapi-generator-cli.jar generate \
	-c config-go.yaml \
    -i openapi.yaml \
    -g go \
    -o go

	cd go && go mod tidy

.PHONY: generate-clients
generate-clients: generate-go-client

.PHONY: golangci-lint
golangci-lint:
	docker run \
		--rm \
		--volume "$(shell pwd):/src" \
		--workdir "/src" \
		golangci/golangci-lint:v1.55 golangci-lint run ./... -v
