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

  for input := range howler.ControlInputs() {
    //input := howler.ControlButton26.ID()
    i, _ := controller.GetInput(howler.ControlID(input))
    fmt.Printf("%02d: %s\n", input, i.String())
  }

  return nil
}
