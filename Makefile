build:
	go build -o mytodo

run: build
	./mytodo

test:
	go test ./... -v
