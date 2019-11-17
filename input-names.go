package main

import (
  "fmt"
  "sort"
  howler "github.com/sferris/howler-controller"
)

func Inputs() string {
  var result string

  var keys []int
  for k := range howler.InputNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.InputNames[howler.Inputs(k)])

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}

func LedInputs() string {
  var result string

  var keys []int
  for k := range howler.LedInputNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.LedInputNames[howler.LedInputs(k)])

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}
