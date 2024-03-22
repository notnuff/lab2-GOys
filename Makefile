default: out/example

clean:
	rm -rf out

test: implementation/*.go handler/*.go
	go test ./...

out/example: implementation/implementation.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
