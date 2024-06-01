package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  } else {
    return fmt.Sprint(math.Sqrt(x))
  }
}

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
    return lim
  }
}

func main() {
  {
    sum := 0
    for i := 0; i < 10; i++ {
      sum += i
    }
    fmt.Println(sum)
  }

  {
    sum := 1
    for ; sum < 1000; {
      sum += sum
    }
    fmt.Println(sum)
  }

  {
    sum := 1
    for sum < 1000 {
      sum += sum
    }
    fmt.Println(sum)
  }

  {
    // for {
    //   fmt.Println(1)
    // }
  }

  {
    fmt.Println(sqrt(2), sqrt(-4))
  }

  {
    fmt.Println(
      pow(3, 2, 10),
      pow(3, 3, 20),
    )
  }

  {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("OS X.")
    case "linux":
      fmt.Println("Linux.")
    default:
      fmt.Printf("%s. \n", os)
    }
  }

  {
    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
      fmt.Println("Today.") 
    case today + 1:
      fmt.Println("Tomorrow")
    case today + 2:
      fmt.Println("In two days.")
    default:
      fmt.Println("Too far away")
    }
  }

  {
    t := time.Now()
    switch {
      case t.Hour() < 12:
        fmt.Println("Good morning!")
      case t.Hour() < 17:
        fmt.Println("Good afternoon.")
      default:
        fmt.Println("Good evening.")
    }
  }

  {
    defer fmt.Println("world")
    fmt.Println("hello")
  }

  {
    fmt.Println("conting")

    for i := 0; i < 10; i++ {
      defer fmt.Println(i)
    } 

    fmt.Println("done")
  }
}


