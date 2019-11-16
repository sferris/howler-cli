package main

import (
  "fmt"
  "log"

  "github.com/sferris/howler-controller"
)

func ShowFirmware() error {
  var err error

  if controller == nil {
    log.Printf("Opening howler device: %d\n", device)

    controller, err = howler.OpenDevice(device)
    if err != nil {
      log.Fatal(err.Error())
    }
  }

  fw, _ := controller.GetFWRelease()
  fmt.Printf("Firmware: %d.%d\n", fw.Major, fw.Minor)

  return err
}

