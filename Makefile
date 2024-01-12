.PHONY: pull-schema, generate-client, golangci-lint

openapi-generator-cli.jar:
	wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar -O openapi-generator-cli.jar

pull-schema:
	curl -H "Authorization: token ${GITHUB_TOKEN}" -H "Accept: application/vnd.github.v3.raw" -O -L https://raw.githubusercontent.com/grafana/slo/julienduchesne/auto-document-huma/docs/sources/api/slo-api.yaml

generate-go-client: openapi-generator-cli.jar pull-schema
	GIT_USER_ID=grafana java -jar openapi-generator-cli.jar generate \
	-c config-go.yaml \
    -i slo-api.yaml \
    -g go \
    -o go

	cd go && go mod tidy

golangci-lint:
	docker run \
		--rm \
		--volume "$(shell pwd):/src" \
		--workdir "/src" \
		golangci/golangci-lint:v1.55 golangci-lint run ./... -v
