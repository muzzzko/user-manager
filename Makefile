
export PROJECT_NAME = user-manager
export PROJECT_PATH = github/$(PROJECT_NAME)
export MIGRATION_IMAGE_NAME = migration-goose

#----------------
# building service
#----------------

build-in-docker:
	go build -o /bin/$(PROJECT_NAME) ./cmd/$(PROJECT_NAME)-server/main.go

build-service-image:
	docker build -t $(PROJECT_NAME) \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		--build-arg PROJECT_PATH=$(PROJECT_PATH) \
		--progress plain \
		-f ./build/Dockerfile .

#----------------
# generating service
#----------------

generate-server-in-docker:
	mkdir -p ./internal/generated
	rm -dfR ./internal/generated/server/models
	rm -dfR ./internal/generated/server/restapi/operations
	docker run --rm -it  --user $(shell id -u):$(shell id -g) -e GOPATH=$(shell go env GOPATH):/go -v $(shell pwd):$(shell pwd) -w $(shell pwd) quay.io/goswagger/swagger:v0.30.3 generate server \
	 	--target=./internal/generated/server \
	 	--template-dir=./api/templates \
	 	--main-package=../../../../cmd/$(PROJECT_NAME)-server \
	 	--exclude-main \
	 	-f ./api/swagger.yml

generate-client-in-docker:
	rm -dfR ./pkg/client
	mkdir -p ./pkg/client
	docker run --rm -it  --user $(shell id -u):$(shell id -g) -e GOPATH=$(shell go env GOPATH):/go -v $(shell pwd):$(shell pwd) -w $(shell pwd) quay.io/goswagger/swagger:v0.30.3 generate client \
	 	--target=./pkg/client \
	 	-A $(PROJECT_NAME) \
	 	-f ./api/swagger.yml

#----------------
# running service
#----------------

run-in-docker:
	docker-compose -f ./build/dev/docker-compose.yml up -d

stop-in-docker:
	docker-compose -f ./build/dev/docker-compose.yml stop

#----------------
# migration
#----------------

build-migration-image:
	docker build -t $(MIGRATION_IMAGE_NAME) \
		--progress plain \
		-f ./build/migration/Dockerfile .

migrate:
	goose --dir ./migration/postgres postgres "$(USER_MANAGER_POSTGRES_MASTER)" up

#----------------
# testing
#----------------

run-test-env-in-docker:
	docker-compose -f ./build/test/docker-compose.yml up -d

stop-test-env-in-docker:
	docker-compose -f ./build/test/docker-compose.yml stop

run-test:
	go test -v ./test/...

run-test-in-docker:
	docker-compose -f ./build/test/docker-compose.yml up test-runner
