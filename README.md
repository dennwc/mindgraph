# MindGraph

Mind mapping with graphs.

## Installation

Compiling the app requires at least Go 1.11 beta1 (`snap install --beta --classic go`)
for WebAssembly support.

```bash
go get -u github.com/dennwc/dom
go install github.com/dennwc/dom/cmd/wasm-server

go get -u github.com/dennwc/mindgraph
cd $GOPATH/src/github.com/dennwc/mindgraph
wasm-server
```

The server should be running on http://localhost:8080.
