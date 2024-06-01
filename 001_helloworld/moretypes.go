package main

import (
  "fmt"
)

type Vertex struct {
  X int
  Y int
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
  {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p)

    p = &j
    *p = *p / 37
    fmt.Println(j)
  }  

  {
    fmt.Println(Vertex{1, 2})
  }

  {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
  }

  {
    v := Vertex{1, 2}
    p := &v
    p.X = 1e9
    fmt.Println(v)
  }

  {
    v1 := Vertex{1, 2}
    v2 := Vertex{X: 1}
    v3 := Vertex{}
    p := &Vertex{1, 2}
    fmt.Println(v1, p, v2, v3)
  }

  {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
  }

  {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)
  }

  {
    names := [4]string{
      "John",
      "Paul",
      "George",
      "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
  }

  {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    s := []struct {
      i int
      b bool
    }{
      {2, true},
      {3, false},
      {5, true},
      {7, true},
      {11, false},
      {13, true},
    }
    fmt.Println(s)
  }

  {
    s := []int{2, 3, 5, 7, 11, 13}

    s = s[1:4]
    fmt.Println(s)

    s = s[:2]
    fmt.Println(s)

    s = s[1:]
    fmt.Println(s)
  }

  {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s) 

    s = s[:0]
    printSlice(s)

    s = s[:4]
    printSlice(s)

    s = s[2:]
    printSlice(s)
  }
}
