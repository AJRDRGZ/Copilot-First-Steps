.PHONY: run test bench fmt vet

run:
	go run ./cmd/katanasvc

test:
	go test ./...

bench:
	go test -bench=. -run=^$

fmt:
	go fmt ./...

vet:
	go vet ./...
