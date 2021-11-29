GO_FMT       = gofmt -s -w -l .
BUILD_TIME   = $(shell date -u +%FT%T)
PKGROOT      = $(shell go list -e 2>/dev/null || true)
SERVICE_NAME = ms-sample
COMMAND_NAME = ms-sample

VERSION     ?= $(shell date -u +%y.%m.%d.%H%M%S)
IMAGE_TAG   ?= $(VERSION)
IMAGE_NAME   = registry.bidmylisting.io/service/$(SERVICE_NAME):$(IMAGE_TAG)
APIS         = $(wildcard api/*/.)
BUILD_TAG   ?= $(SERVICE_NAME)

run: build-local
#	@echo starting database container
#	@docker-compose up -d db
	@./bin/$(COMMAND_NAME)
	@if [ -f ./local.yml ]; then ./bin/$(COMMAND_NAME) -c ./local.yml; else ./bin/$(COMMAND_NAME); fi

all: check test build-local

check: format vet

format:
	@echo "formatting files..."
	@echo "  > getting goimports" && go install golang.org/x/tools/cmd/goimports@latest
	@echo "  > getting goimports-blank-rm" && go install github.com/jucardi/goimports-blank-rm@latest
	@echo "  > executing goimports-blank-rm" && goimports-blank-rm . 1>/dev/null 2>/dev/null
	@echo "  > executing goimports" && goimports -w -l $(shell find . -type f -name '*.go' -not -path "./vendor/*") 1>/dev/null 2>/dev/null || true
	@echo "  > executing gofmt" && gofmt -s -w -l $(shell find . -type f -name '*.go' -not -path "./vendor/*") 1>/dev/null

vet:
	@echo "vetting..."
	@go vet -mod=vendor ./...

deps: templates
	@echo "installing dependencies..."
	@[ -f ./go.mod ] || go mod init
	@go mod tidy
	@go mod vendor

templates: protoc format

protoc: protoc-dep protoc-item

protoc-dep:
	@echo "getting protobuf dependencies..."
	@go get google.golang.org/protobuf/protoc-gen-go
	@go install github.com/jucardi/protoc-go-inject-tag@latest 2>/dev/null 1>/dev/null

protoc-item: $(APIS)

$(APIS):
	@[ -f $@/*.proto ] && echo "$@  > generating protobuf..." || echo "$@  > no protobuf files found, skipping."
	@[ -f $@/*.proto ] && protoc -I=$@ --go_out=$@ --proto_path=$@ $@/*.proto || true
	@[ -f $@/*.proto ] && echo "$@  > adding golang tags..." || true
	@[ -f $@/*.proto ] && protoc-go-inject-tag --input "$@/*.pb.go" --cleanup -x yaml -x gorm -x bson -x structs 2>/dev/null || true

.PHONY: protoc-item $(APIS)

test:
	@echo "running test coverage..."
	@mkdir -p test-artifacts/coverage
	@go test -p 1 -mod=vendor ./... -v -coverprofile test-artifacts/cover.out
	@go tool cover -func test-artifacts/cover.out

build-local:
	@go build -mod=vendor -ldflags "-X $(COMMONS_ROOT)/info.Version=$(VERSION) -X $(COMMONS_ROOT)/info.Built=$(BUILD_TIME)" -o ./bin/$(COMMAND_NAME) ./cmd/service

build-linux:
	@mkdir bin 2>/dev/null || true
	@echo "building linux binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags "-X $(COMMONS_ROOT)/info.Version=$(VERSION) -X $(COMMONS_ROOT)/info.Built=$(BUILD_TIME)" -o bin/$(COMMAND_NAME)-linux ./cmd/server

docker-cleanup:
	@VERSION=$(VERSION) docker-compose down 1>/dev/null
	@VERSION=$(VERSION) docker-compose -p $(BUILD_TAG) -f build-compose.yml -f docker-compose.yml down 1>/dev/null

docker-run: docker-cleanup build-linux
	@docker network create global 2>/dev/null || true
	@VERSION=$(VERSION) docker-compose up --force-recreate --build --abort-on-container-exit --exit-code-from api api

docker-build: docker-cleanup
	@VERSION=$(VERSION) docker-compose -p $(BUILD_TAG) -f build-compose.yml up --abort-on-container-exit --exit-code-from builder builder

docker-image:
	@docker build -t $(IMAGE_NAME) .
	@echo $(IMAGE_NAME) > .docker_image_name

docker-image-local: build-linux
	@docker build -t $(IMAGE_NAME):latest .

dep-update-titan:
	@go get -u github.com/jucardi/go-titan