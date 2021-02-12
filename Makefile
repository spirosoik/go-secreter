.PHONY: lint setup test

lint: setup ## Linting the codebase
	golint -set_exit_status ./...

run-example: ## Run the monolith API
	go run ./example/

setup: ## Setup modules
	go get -u golang.org/x/lint/golint

test: ## Run all tests
	go test -coverprofile=cover.out -coverpkg=$(go list ./...)  ./...
	go tool cover -func=cover.out

race: ## Run data race detector
	go test -race -short ./...

clean: ## Clean
	go clean
	rm -f cover.out