# Wasmer sample

This sample shows how to run WebAssembly inside EGo using [Wasmtime](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go).

By default, *wasmer-go* comes with a shared library. EGo only supports static linking. To this end, download the wasmer static library and tell the Go compiler to use it:
```sh
wget -O- https://github.com/wasmerio/wasmer/releases/download/2.2.1/wasmer-linux-amd64.tar.gz | tar xz --one-top-level=wasmer
CGO_CFLAGS="-I$PWD/wasmer/include" CGO_LDFLAGS="$PWD/wasmer/lib/libwasmer.a -ldl -lm -static-libgcc" ego-go build -tags custom_wasmer_runtime
```

Then you can sign and run as usual:
```sh
ego sign wastime_sample
ego run wastime_sample
```

You should see an output similar to:
```
[erthost] loading enclave ...
[erthost] entering enclave ...
[ego] starting application ...
Results of `sum`: 3
```

Note that `executableHeap` is enabled in `enclave.json` so that Wasmer can JIT-compile the WebAssembly.
