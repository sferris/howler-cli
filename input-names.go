package main

import (
  "fmt"
  "sort"
  howler "github.com/sferris/howler-controller"
)

func ControlInputs() string {
  var result string

  var keys []int
  for k := range howler.ButtonNames {
    keys = append(keys, int(k))
  }
  for k := range howler.JoystickNames {
    keys = append(keys, int(k))
  }
  for k := range howler.AxisNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.ControlInput(k))

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
    value := fmt.Sprintf("%s, ", howler.LedInputs(k))

    result += value

    w += len(value)
    if w >= (columns-20) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result + fmt.Sprintln()
}
