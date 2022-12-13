OUTPUT := server.out
DYNAMO_ENDPOINT := http://localhost:8000
API_PORT := 80

start-vm: ## Start vm
	finch vm start

setup: ## Set up container
	mkdir -p docker/dynamodb
	finch compose up -d

create-tables: ## Create dynamodb tables
	aws dynamodb create-table --table-name Game \
        --attribute-definitions AttributeName=Id,AttributeType=S \
        --key-schema AttributeName=Id,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
        --endpoint-url $(DYNAMO_ENDPOINT)

drop-tables: ## Drop dynamodb tables
	aws dynamodb delete-table \
		--table-name Game \
        --endpoint-url $(DYNAMO_ENDPOINT)

drop-create-tables: drop-tables create-tables ## Drop & Create dynamodb tables

scan-dynamodb-game: ## Scan Game table
	@aws dynamodb scan \
		--table-name Game \
        --endpoint-url $(DYNAMO_ENDPOINT)

describe-dynamodb-game: ## Describe Game table define
	aws dynamodb \
		describe-table \
		--endpoint-url $(DYNAMO_ENDPOINT) \
		--table-name Game

build: ## Build files
	go build -o $(OUTPUT) cmd/server/main.go

run: build ## Build & Run
	PORT=$(API_PORT) ./$(OUTPUT)

fmt: ## Format go files
	gofmt -l -s -w .

test: ## Execute test
	go test -v ./...

clean: ## Clean files
	rm $(OUTPUT)
	go clean

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
