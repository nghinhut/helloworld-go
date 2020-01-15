package main

import "fmt"

func numbersPrinter() {
  numberArray := []int32{7, 1, 2, 5, 8, 10, 6, 4}

  fmt.Println("INDEX | VALUE")
  for index, value := range numberArray {
    fmt.Printf("%d \t\t %d\n", index, value)
  }

}

