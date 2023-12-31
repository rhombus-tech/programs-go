package main

import (
	"fmt"

	"github.com/bytecodealliance/wasmtime-go/v12"
)

// https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go#example-Config-Fuel
func main() {
	// Let's assume we don't have WebAssembly bytes at hand. We
	// will write WebAssembly manually.
	wasmBytes := []byte(`
	(module
	  (type (func (param i32 i32) (result i32)))
	  (func (type 0)
	    local.get 0
	    local.get 1
	    i32.add)
	  (export "sum" (func 0)))
`)

    cfg := wasmtime.NewConfig()
    cfg.SetConsumeFuel(true)
	cfg.CacheConfigLoadDefault()
	cfg.SetStrategy(wasmtime.StrategyCranelift)
	store := wasmtime.NewStore(wasmtime.NewEngineWithConfig(cfg))

	err := store.AddFuel(10000000000000) // testing only ;)
	if err != nil {
		fmt.Println("Failed to add fuel:", err)
	}

	module, err := wasmtime.NewModule(store.Engine, wasmBytes)
	if err != nil {
		fmt.Println("Failed to compile module:", err)
	}

	// Create an empty import object.
	importObject := wasmer.NewImportObject()

	// Let's instantiate the WebAssembly module.
	instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		panic(fmt.Sprintln("Failed to instantiate the module:", err))
	}

	// // Now let's execute the `sum` function.
	// sum, err := instance.Exports.GetFunction("sum")

	// if err != nil {
	// 	panic(fmt.Sprintln("Failed to get the `add_one` function:", err))
	// }

	// result, err := sum(1, 2)

	// if err != nil {
	// 	panic(fmt.Sprintln("Failed to call the `add_one` function:", err))
	// }

	// fmt.Println("Results of `sum`:", result)
}
