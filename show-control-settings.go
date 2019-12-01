package main

import (
  "fmt"
  _"strings"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

func showControlSettings(c *cli.Context) error {
  var err error

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  for i, control := range howler.ControlInputsByID() {
    input, err := controller.GetInput(control.ID())
    if err != nil { 
      fmt.Printf("%02d %+v\n", i, input.Dump())
    } else {
      fmt.Printf("%02d %s\n", i, input.String())
    }
  }

  return nil
}
