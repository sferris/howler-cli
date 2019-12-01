package main

import (
  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

func resetDefaults(c *cli.Context) error {
  var err error

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  _, err = controller.ResetToDefaults()
  return err
}
