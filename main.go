package main

import (
	"context"
	"fmt"

	"github.com/ava-labs/hypersdk/x/programs/examples/imports/program"
	"github.com/ava-labs/hypersdk/x/programs/examples/imports/pstate"

	"github.com/ava-labs/hypersdk/x/programs/runtime"
)

func main() {
	// Get a greeting message and print it.
	db := newTestDB()
	var counterProgramBytes []byte
	fmt.Println("hello world")
	maxUnits := uint64(80000)
	cfg, err := runtime.NewConfigBuilder().Build()
	fmt.Println(err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define supported imports
	supported := runtime.NewSupportedImports()
	supported.Register("state", func() runtime.Import {
		return pstate.New(log, db)
	})
	supported.Register("program", func() runtime.Import {
		return program.New(log, db, cfg)
	})

	rt := runtime.New(log, cfg, supported.Imports())
	err = rt.Initialize(ctx, counterProgramBytes, maxUnits)
	rt.Stop()
}
