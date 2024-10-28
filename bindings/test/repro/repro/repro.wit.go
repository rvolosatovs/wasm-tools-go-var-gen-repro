// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package repro represents the imported interface "test:repro/repro".
package repro

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// Var represents the variant "test:repro/repro#var".
//
//	variant var {
//		var(u8),
//	}
type Var cm.Variant[uint8, uint8, uint8]

// VarVar_ returns a [Var] of case "var".
func VarVar_(data uint8) Var {
	return cm.New[Var](0, data)
}

// Var_ returns a non-nil *[uint8] if [Var] represents the variant case "var".
func (self *Var) Var_() *uint8 {
	return cm.Case[uint8](self, 0)
}

var stringsVar = [1]string{
	"var",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v Var) String() string {
	return stringsVar[v.Tag()]
}

// Foo represents the imported function "foo".
//
//	foo: func(var: var)
//
//go:nosplit
func Foo(var_ Var) {
	var0, var1 := lower_Var(var_)
	wasmimport_Foo((uint32)(var0), (uint32)(var1))
	return
}