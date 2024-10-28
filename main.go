//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world imports --out bindings ./wit

package main

import _ "github.com/rvolosatovs/wasm-tools-go-var-gen-repro/bindings/test/repro/repro"

func init() {
}

func main() {}
