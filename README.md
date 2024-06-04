```bash
$ cp /usr/lib/go-1.22/misc/wasm/wasm_exec.js .
$ go get github.com/junegunn/fzf
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ deno run --allow-read main.ts
```

