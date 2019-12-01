package main

import (
  "fmt"
  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

func showFirmwareRevision(c *cli.Context) error {
  var err error

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  fw, _ := controller.GetFWRelease()
  fmt.Printf("Firmware: %d.%d\n", fw.Major, fw.Minor)

  return err
}

