package main

import (
  // "fmt"
  // "github.com/junegunn/fzf/src/algo"
  // "github.com/junegunn/fzf/src/util"
  "syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
  a := args[0].Int()
  b := args[1].Int()
  return js.ValueOf(a + b)
}

func main() {
  c := make(chan struct{}, 0)
  js.Global().Set("add", js.FuncOf(add))
  <-c
  // input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  // pattern := []rune("o")
  // slab := util.MakeSlab(100*1024, 2048)
  // result, _ := algo.FuzzyMatchV2(false, true, true, &input, pattern, false, slab)
  // fmt.Printf("Result: %+v\n", result)
}
