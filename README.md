# Learning WebAssembly with Golang

Initial setup:
1. Copy WASM loader: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
2. Create [index.html](index.html)

Development flow:

1. edit golang code
2. `GOOS=js GOARCH=wasm go build -o main.wasm`
3. `goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'`
4. open browser - see console output


See: https://github.com/stdiopt/gowasm-experiments