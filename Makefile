PKGS=$(shell go list ./...)

lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.39.0; \
	fi
	golangci-lint run

test:
	go test -race -cover -v -count=1 $(PKGS)

cover:
	go test -race -v -count=1 -coverprofile=coverage.txt $(PKGS)
	go tool cover -func=coverage.txt

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/app

.PHONY: lint test build cover
