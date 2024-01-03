# Wasmer sample

This sample shows how to run WebAssembly inside EGo using [Wasmtime](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go).

By default, *wasmtime-go* comes with a shared library. EGo only supports static linking. To this end, download the wasmtime static library and tell the Go compiler to use it:
```sh
wget -O- https://github.com/bytecodealliance/wasmtime/releases/download/v15.0.1/wasmtime-v12.0.0-x86_64-linux-c-api.tar.xz  | tar xf --one-top-level=wasmer
CGO_CFLAGS="-I$PWD/wasmtime/include" CGO_LDFLAGS="$PWD/wasmtime/lib/libwasmtime.a -ldl -lm -static-libgcc" ego-go build -tags custom_wasmtime_runtime
```

Then you can sign and run as usual:
```sh
ego sign wasmtime_sample
ego run wasmtime_sample
```

You should see an output similar to:
```
[erthost] loading enclave ...
[erthost] entering enclave ...
[ego] starting application ...
Results of `sum`: 3
```

Note that `executableHeap` is enabled in `enclave.json` so that Wasmer can JIT-compile the WebAssembly.
