package main

import (
  "os"
  "fmt"
  _"time"

  _"encoding/hex"

  howler "github.com/sferris/howler-controller"
)

func main() {

/*
  device := howler.HowlerConfig{}
*/

  device, err := howler.OpenHowlerConfig(0)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  button := howler.Button26

/*
  fmt.Println("Write")
  var data = []byte{0xce,0x01,button,0xff,0x00,0xff,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  device.Write(data)

  fmt.Println("Query")
  var stmt = []byte{0xce,0x08,button,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  device.WriteWithResponse(stmt)
*/

  device.SetDefaultLEDColor(button,"orange")

  device.SetLEDColor(button,"red")

  device.SetLEDColor(button,"green")

  device.SetLEDColor(button,"blue")

  for i := 0; i < int(howler.ButtonMax); i++ {
    c, _ := device.GetLEDColor(howler.Buttons(i))
    fmt.Printf("%d: R: %d, G: %d, B: %d\n", i, c.R, c.G, c.B)
  }

  fw, _ := device.GetFWRelease()
  fmt.Printf("Firmware: %s\n", fw)

  for i := 0; i < int(howler.InputMax); i++ {
    err := device.GetInput(howler.Inputs(i))
    if err != nil {
      fmt.Println(err.Error())
    }
  }

/*
 device.SetInput(howler.InputButton26, howler.ModeJoystick2, 13, howler.ModifierNone)
 time.Sleep(time.Millisecond*500)

 device.GetInput(howler.InputButton26)
*/
}
