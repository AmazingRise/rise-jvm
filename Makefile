main.wasm: main.go
	tinygo build -o main.wasm -target wasm ./main.go
	#GOARCH=wasm GOOS=js go build -o main.wasm main.go

wasm_exec.js:
	cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js .
	#cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" .

.PHONY: clean
clean:
	rm -rf main.wasm wasm_exec.js

.PHONY: all
all: main.wasm wasm_exec.js

