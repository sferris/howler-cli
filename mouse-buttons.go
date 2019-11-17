package main

import (
  "fmt"
  "sort"
  howler "github.com/sferris/howler-controller"
)

func MouseButtons() string {
  var result string

  var keys []int
  for k := range howler.MouseButtonNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.MouseButtonNames[howler.MouseButtons(k)])

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}
