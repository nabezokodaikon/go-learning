package main

import (
  "encoding/json"
  // "fmt"
  // "github.com/junegunn/fzf/src/algo"
  // "github.com/junegunn/fzf/src/util"
  // "strconv"
  "syscall/js"
)

type Animal struct {
  Name string
  Hoge int
}

func find(this js.Value, args []js.Value) interface{} {
  length := args[1].Int()
  jsArray := js.Global().Get("Array").New(length)
  for i := 0; i < length; i++ {
    v := args[i + 2];
    str := js.Global().Get("JSON").Call("stringify", v).String()
    var dst Animal
    json.Unmarshal([]byte(str), &dst)
    jsObj := js.Global().Get("Object").New()
    jsObj.Set("Name", dst.Name)
    jsObj.Set("Hoge", 777)
    jsArray.SetIndex(i, jsObj)
  }

  return jsArray
}

func main() {
  c := make(chan struct{}, 0)
  js.Global().Set("find", js.FuncOf(find))
  <-c

  // input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  // pattern := []rune("o")
  // slab := util.MakeSlab(100*1024, 2048)
  // result, _ := algo.FuzzyMatchV2(false, true, true, &input, pattern, false, slab)
  // fmt.Printf("Result: %+v\n", result)
}
