package main

import (
  "os"
  "fmt"
  "time"

  howler "github.com/sferris/howler-controller"
)

func main() {
  fmt.Println("Hello World!")

  device, err := howler.OpenHowlerConfig(0)
  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  device.SetDefaultLEDColor(howler.Button26,"orange")
  time.Sleep(time.Second*2)

  device.SetLEDColor(howler.Button("button26"),"red")
  time.Sleep(time.Second*2)

  device.SetLEDColor(howler.Button("button26"),"green")
  time.Sleep(time.Second*2)

  device.SetLEDColor(howler.Button("button26"),"blue")

/*
  for i := 0; i < int(howler.MaxButton); i++ {
    c, _ := device.GetLEDColor(howler.Button(fmt.Sprintf("%d",i)))
    fmt.Printf("R: %d, G: %d, B: %d\n", c.R, c.G, c.B)
    time.Sleep(time.Millisecond*500)
  }
*/

  fw, _ := device.GetFWRelease()
  fmt.Printf("Firmware: %s\n", fw)

/*
  for i := 0; i < int(howler.MaxButton); i++ {
    _, err := device.GetInput(howler.Input(fmt.Sprintf("%d",i)))
    if err != nil {
      fmt.Println(err.Error())
    }
    time.Sleep(time.Millisecond*500)
  }
*/

 device.SetInput(howler.InputButton26, howler.ModeJoystick2, 13, howler.ModifierNone)
 time.Sleep(time.Millisecond*500)

 device.GetInput(howler.InputButton26)
}
