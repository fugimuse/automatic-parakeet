.PHONY: run build test vet fmt tidy clean debug upgrade

run:
	go run .

build:
	go build -o bin/budget-app .

test:
	go test ./...

vet:
	go vet ./...

fmt:
	goimports -w .

tidy:
	go mod tidy

clean:
	rm -rf bin/
	find . -type d -name ".approval_tests_temp" -exec rm -rf {} +

debug:
	dlv debug --headless --api-version=2 --listen=127.0.0.1:43000 .

upgrade:
	go get -u ./...
	go mod tidy
