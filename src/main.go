package main

import (
  "encoding/json"
  // "fmt"
  "golang.org/x/exp/maps"
  // "github.com/junegunn/fzf/src/algo"
  // "github.com/junegunn/fzf/src/util"
  // "strconv"
  "strconv"
  "syscall/js"
)

type Animal struct {
  Name string
  Hoge int
}

type Record struct {
  int
  string
}

func structure(this js.Value, args []js.Value) interface{} {
  length := args[0].Int()
  jsArray := js.Global().Get("Array").New(length)
  for i := 0; i < length; i++ {
    v := args[i + 1];
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

func record(this js.Value, args []js.Value) interface{} {
  v := args[0]
  str := js.Global().Get("JSON").Call("stringify", v).String()
  var dst map[int]string
  json.Unmarshal([]byte(str), &dst)
  keys := maps.Keys(dst)
  key := int(keys[0])
  jsObj := js.Global().Get("Object").New()
  jsObj.Set(strconv.Itoa(key), dst[key])
  return jsObj
}

func main() {
  c := make(chan struct{}, 0)
  // js.Global().Set("structure", js.FuncOf(structure))
  js.Global().Set("record", js.FuncOf(record))
  <-c

  // input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  // pattern := []rune("o")
  // slab := util.MakeSlab(100*1024, 2048)
  // result, _ := algo.FuzzyMatchV2(false, true, true, &input, pattern, false, slab)
  // fmt.Printf("Result: %+v\n", result)
}
