package main

import (
  "fmt"
  "github.com/junegunn/fzf/src/algo"
  "github.com/junegunn/fzf/src/util"
)

func main() {
  input := util.RunesToChars([]rune("github.com/wantedly/wantedly"))
  pattern := []rune("o")
  slab := util.MakeSlab(100*1024, 2048)
  result, _ := algo.FuzzyMatchV2(false, true, true, &input, pattern, false, slab)
  fmt.Printf("Result: %+v\n", result)
}
