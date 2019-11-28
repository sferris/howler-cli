package main

import (
  "fmt"
  "sort"
  howler "github.com/sferris/howler-controller"
)


type controlSlice []howler.ControlInputs

func (controls controlSlice) Len() int {
  return len(controls)
}

func (controls controlSlice) Swap(i, j int) {
  controls[i], controls[j] = controls[j], controls[i]
}

func (controls controlSlice) Less(i, j int) bool {
  return int(controls[i].Input()) < int(controls[j].Input())
}

func ControlInputs() string {
  var result string

  controls := make(controlSlice,0,len(howler.ControlInputNames))
  for _, control := range howler.ControlInputNames {
    controls = append(controls, control)
  }

  sort.Sort(controls)

  w := 0
  for _, v := range controls {
    value := fmt.Sprintf("%s, ", v.Name())

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
