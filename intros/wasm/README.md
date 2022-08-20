# Quick And Easy WASM

1. `syscall/js` is necessary
1. `GOOS=js GOARCH=wasm tinygo build -o static/main.wasm cmd/wasm/main.go`
1. `cp "$(tinygo env TINYGOROOT)"/targes/wasm/wasm_exec.js ./static`
1. `<head><script src="wasm_exec.js"><script> const go = new Go() ...`
1. Call named Go functions in js `wasmHash(inputField.value)`

## Debugging

https://blog.noops.land/debugging-webAssembly-from-go-sources-in-chrome-devtools
