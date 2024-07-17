#!/bin/bash
set -eo pipefail

if [ ! -f openapi-generator-cli.jar ]; then
  # Download openapi-generator-cli if not present
  wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.2.0/openapi-generator-cli-7.2.0.jar -O ./openapi-generator-cli.jar
else
  echo "openapi-generator-cli.jar already exists. Delete it to download a new version."
fi

# Cleanup old files
rm -rf ./go

java -jar openapi-generator-cli.jar generate \
  -i openapi.yaml \
  -g go \
  -o go \
  -c ./scripts/config-go.yaml \
  -p disallowAdditionalPropertiesIfNotPresent=false \
  -t ./scripts

pushd go && go mod tidy && popd

if ! command -v goimports &> /dev/null
then
    go install golang.org/x/tools/cmd/goimports@latest
fi
find ./go -name \*.go -exec goimports -w {} \;

