.PHONY: run
run:
	go run --race cmd/api/main.go

.PHONY: tidy
tidy:
	go mod tidy -v
