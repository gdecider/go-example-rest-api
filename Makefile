.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: migrate-up
migrate-up:
	migrate -path migrations -database "mysql://root:@tcp(127.0.0.1:3306)/gorestapi" up

.PHONY: migrate-down
migrate-down:
	migrate -path migrations -database "mysql://root:@tcp(127.0.0.1:3306)/gorestapi" down

.PHONY: migrate-up
migrate-test-up:
	migrate -path migrations -database "mysql://root:@tcp(127.0.0.1:3306)/gorestapi_test" up

.PHONY: migrate-test-down
migrate-down:
	migrate -path migrations -database "mysql://root:@tcp(127.0.0.1:3306)/gorestapi_test" down


.DEFAULT_GOAL := build
