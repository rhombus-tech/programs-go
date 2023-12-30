package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"github.com/ava-labs/hypersdk/x/programs/examples/imports/program"
	"github.com/ava-labs/hypersdk/x/programs/examples/imports/pstate"

	"github.com/ava-labs/avalanchego/utils/logging"

	"github.com/ava-labs/hypersdk/x/programs/runtime"
	"github.com/ava-labs/avalanchego/database/memdb"
)

var (
	//go:embed testdata/token.wasm
	tokenProgramBytes []byte

	log = logging.NewLogger(
		"",
		logging.NewWrappedCore(
			logging.Info,
			os.Stderr,
			logging.Plain.ConsoleEncoder(),
		))
)

type testDB struct {
	db *memdb.Database
}

func newTestDB() *testDB {
	return &testDB{
		db: memdb.New(),
	}
}

func (c *testDB) GetValue(_ context.Context, key []byte) ([]byte, error) {
	return c.db.Get(key)
}

func (c *testDB) Insert(_ context.Context, key []byte, value []byte) error {
	return c.db.Put(key, value)
}

func (c *testDB) Put(key []byte, value []byte) error {
	return c.db.Put(key, value)
}

func (c *testDB) Remove(_ context.Context, key []byte) error {
	return c.db.Delete(key)
}

func main() {
	// Get a greeting message and print it.
	db := newTestDB()
	var counterProgramBytes []byte
	fmt.Println("hello world")
	maxUnits := uint64(80000)
	cfg, err := runtime.NewConfigBuilder(maxUnits).Build()
	fmt.Println(err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// define supported imports
	supported := runtime.NewSupportedImports()
	supported.Register("state", func() runtime.Import {
		return pstate.New(log, db)
	})
	supported.Register("program", func() runtime.Import {
		return program.New(log, db)
	})

	rt := runtime.New(log, cfg, supported.Imports())
	err = rt.Initialize(ctx, counterProgramBytes)
	rt.Stop()
}
