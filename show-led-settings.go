
package main

import (
  "fmt"
  _"strings"

  "gopkg.in/urfave/cli.v2"

  "github.com/sferris/howler-controller"
)

func showLEDSettings(c *cli.Context) error {
  for led := howler.LedMin; led < howler.LedMax; led++ {
    var err error

    if controller == nil {
      controller, err = howler.OpenDevice(device)
      if err != nil {
        return err
      }
    }

    c, _ := controller.GetLEDColor(howler.LedInputs(led))
    fmt.Printf("Led %-15s %s\n", 
      fmt.Sprintf("%s:", howler.LedInputs(led)), c.ToIntString())
  }

  return nil
}
