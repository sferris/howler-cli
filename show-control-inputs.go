package main

import (
  "fmt"
  "sort"
  "gopkg.in/urfave/cli.v2"
  howler "github.com/sferris/howler-controller"
)

type controlSlice []howler.ControlInput

func (controls controlSlice) Len() int {
  return len(controls)
}

func (controls controlSlice) Swap(i, j int) {
  controls[i], controls[j] = controls[j], controls[i]
}

func (controls controlSlice) Less(i, j int) bool {
  return int(controls[i].ID()) < int(controls[j].ID())
}

func showControlInputs(c *cli.Context) error {
  fmt.Println("Valid Control inputs:\n")
  fmt.Println( ControlInputs() );

  return nil
}

func ControlInputs() string {
  var result string

  controls := make(controlSlice,0,len(howler.ControlInputMap))
  for _, control := range howler.ControlInputMap {
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
