build: build-eval build-moves
build-eval:
	go build -o bin/evaluate cmd/evaluate/main.go
build-moves:
	go build -o bin/valid-moves cmd/valid-moves/main.go

build-wasm:
	GOOS=js GOARCH=wasm go build -o bin/evaluate.wasm cmd/evaluate/main.go
	GOOS=js GOARCH=wasm go build -o bin/valid-moves.wasm cmd/valid-moves/main.go

test:
	go test ./...
