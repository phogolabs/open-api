#!/bin/bash

rm -fr ./target
rm -fr ./test

mvn package

chmod +x ./target//golang-server-openapi-generator-1.0.0.jar
# export GO_POST_PROCESS_FILE="go fmt -w"

  # -DdebugOpenAPI \
  # -DdebugModels \
  # -DdebugOperations \
  # -DdebugSupportingFiles \
  # -i https://api.openbankproject.com/obp/v3.1.0/resource-docs/v3.0.0/swagger \
java \
  -cp /usr/local/Cellar/openapi-generator/3.3.4/libexec/openapi-generator-cli-3.3.4.jar:./target/golang-server-openapi-generator-1.0.0.jar org.openapitools.codegen.OpenAPIGenerator \
  generate -g golang-server \
  -i ./petstore.yaml \
  -o ./test \
  --enable-post-process-file \
  --git-user-id phogolabs \
  --git-repo-id pet-store
