package main
import (
  "encoding/json"
  "fmt"
  // "golang.org/x/exp/maps"
  "github.com/junegunn/fzf/src/algo"
  "github.com/junegunn/fzf/src/util"
  // "strconv"
  "syscall/js"
)

type ItemHighlight struct {
  Name string
  Hl_group string
  Col int
  Width int
}

type ItemStatus struct {
  Size int
  Time int
}

type DduItem struct {
  Word string
  Display string
  // action ActionData
  // data unknown
  Highlights []ItemHighlight
  Status ItemStatus
  Kind string
  Level int
  TreePath string
  IsExpanded bool 
  IsTree bool

  MatcherKey string
  SourceIndex int `json:"__sourceIndex"`
  SourceName string `json:"__sourceName"`
  Level_ int `json:"__level"`
  Expanded bool `json:"__expanded"`
  ColumnTexts map[int]string `json:"__columnTexts"`
  GroupedPath string `json:"__groupedPath"`
}

type FindResult struct {
  Score int
  Positions []int
  Item DduItem
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

func toJson(findResult FindResult) interface{} {
  return map[string]interface{}{
    "user": findResult.Item.Word,
  }
}

func find(this js.Value, args []js.Value) interface{} {
  input := args[0]
  fmt.Println(input)
  length := args[1].Int()
  fmt.Println(length)
  jsArray := js.Global().Get("Array").New(length)
  fmt.Println(jsArray)
  var jsonSlice []interface{}
  fmt.Println(jsonSlice)
  for i := 0; i < length; i++ {
    v := args[i + 2]
    str := js.Global().Get("JSON").Call("stringify", v).String()
    fmt.Println(str)
    var item DduItem
    json.Unmarshal([]byte(str), &item)
    fmt.Println(item)

    word := util.RunesToChars([]rune(item.Word))
    util.RunesToChars([]rune(item.Word))
    inputRune := []rune(input)
    slab := util.MakeSlab(100*1024, 2048)
    result, positions := algo.FuzzyMatchV2(false, true, true, &word, inputRune, true, slab)
    findResult := FindResult {
      Score: result.Score,
      Positions: *positions,
      Item: item,
    }
    findResultJson := toJson(findResult)
    jsonSlice = append(jsonSlice, findResultJson) 


    // jsObj.Set("Name", dst.Name)
    // jsObj.Set("Hoge", 777)
    // jsArray.SetIndex(i, jsObj)
  }

  return nil
}

func main() {
  c := make(chan struct{}, 0)
  js.Global().Set("find", js.FuncOf(find))
  <-c

  // input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  // pattern := []rune("gi")
  // slab := util.MakeSlab(100*1024, 2048)
  // result, positons := algo.FuzzyMatchV2(false, true, true, &input, pattern, true, slab)
  // fmt.Printf("Result: %+v\n", result)
  // fmt.Println(positons)
}
