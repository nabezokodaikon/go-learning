package main

import (
  "encoding/json"
  "fmt"
  // "golang.org/x/exp/maps"
  // "github.com/junegunn/fzf/src/algo"
  // "github.com/junegunn/fzf/src/util"
  // "strconv"
  "syscall/js"
)

type Animal struct {
  Name string
  Hoge int
}

type ItemHighlight struct {
  name string
  hl_group string
  col int
  width int
}

type ItemStatus struct {
  size int
  time int
}

type Record struct {
  int
  string
}

type DduItem struct {
  word string
  display string
  // action ActionData
  // data unknown
  highlights []ItemHighlight
  status ItemStatus
  kind string
  level int
  treePath string
  isExpanded bool 
  isTree bool

  matcherKey string
  __sourceIndex int
  __sourceName string
  __level int
  __expanded bool
  __columnTexts Record
  __groupedPath string
}

type FindResult struct {
  Score int
  Positions []int
  item DduItem
}

// func structure(this js.Value, args []js.Value) interface{} {
//   length := args[0].Int()
//   jsArray := js.Global().Get("Array").New(length)
//   for i := 0; i < length; i++ {
//     v := args[i + 1];
//     str := js.Global().Get("JSON").Call("stringify", v).String()
//     var dst Animal
//     json.Unmarshal([]byte(str), &dst)
//     jsObj := js.Global().Get("Object").New()
//     jsObj.Set("Name", dst.Name)
//     jsObj.Set("Hoge", 777)
//     jsArray.SetIndex(i, jsObj)
//   }
//
//   return jsArray
// }
//
// func record(this js.Value, args []js.Value) interface{} {
//   v := args[0]
//   str := js.Global().Get("JSON").Call("stringify", v).String()
//   var dst map[int]string
//   json.Unmarshal([]byte(str), &dst)
//   keys := maps.Keys(dst)
//   key := int(keys[0])
//   jsObj := js.Global().Get("Object").New()
//   jsObj.Set(strconv.Itoa(key), dst[key])
//   return jsObj
// }

func find(this js.Value, args []js.Value) interface{} {
  input := args[0]
  fmt.Println(input)
  length := args[1].Int()
  // jsArray := js.Global().Get("Array").New(length)
  for i := 0; i < length; i++ {
    v := args[i + 1];
    str := js.Global().Get("JSON").Call("stringify", v).String()
    var dst DduItem
    json.Unmarshal([]byte(str), &dst)
    fmt.Println(dst)
    // jsObj := js.Global().Get("Object").New()
    // jsObj.Set("Name", dst.Name)
    // jsObj.Set("Hoge", 777)
    // jsArray.SetIndex(i, jsObj)
  }

  return nil
}

func main() {
  c := make(chan struct{}, 0)
  // js.Global().Set("structure", js.FuncOf(structure))
  js.Global().Set("find", js.FuncOf(find))
  <-c

  // input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  // pattern := []rune("gi")
  // slab := util.MakeSlab(100*1024, 2048)
  // result, positons := algo.FuzzyMatchV2(false, true, true, &input, pattern, true, slab)
  // fmt.Printf("Result: %+v\n", result)
  // fmt.Println(positons)
}
