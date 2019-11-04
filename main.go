package main

import (
  "os"
  _"fmt"
  "flag"

  _"time"

  //howler "github.com/sferris/howler-controller"
)

/*
type Input struct {
  Button    string
  Mode      string
  Modifier  string
  Value     string
}

type Led struct {
  Button    string
  Mode      string
  RGB       []int
}
*/


func main() {
  setLedCMD := flag.NewFlagSet("set-led", flag.ExitOnError)
  setInputCMD := flag.NewFlagSet("set-input", flag.ExitOnError)

  LedButton := setLedCMD.String("button", "", "Button to set Led color (Required)")
  LedMode   := setLedCMD.String("mode", "", "The Led mode [Immediate||Default] (Required) ")
  LedRGB    := setLedCMD.String("RGB", "", "The Led mode [Immediate||Default] (Required) ")

  InputButton   := setInputCMD.String("button", "", 
    "Button to set the Input properties (Required)")
  InputMode     := setInputCMD.String("mode", "", 
    "The Input mode [Joystick1 or 2||Keyboard||Mouse] (Required) ")
  InputModifier := setInputCMD.String("modifier", "", 
    "The Led mode [Immediate||Default] (Required) ")
  InputValue    := setInputCMD.String("value", "", 
    "The Led mode [Immediate||Default] (Required) ")

  os.Exit(0)
}

/*
  //howler.DumpDevices()

  device, err := howler.OpenDevice(0)
  defer func() {
    device.Close()
  }()

  if err != nil {
    fmt.Println(err.Error())
    os.Exit(1)
  }

  button := howler.Button26

  fmt.Printf("Got here: %d\n", 1);
  device.SetDefaultLEDColor(button,"orange")
  time.Sleep(time.Second);

  fmt.Printf("Got here: %d\n", 2);
  device.SetLEDColor(button,"red")
  time.Sleep(time.Second);

  fmt.Printf("Got here: %d\n", 3);
  device.SetLEDColor(button,"green")
  time.Sleep(time.Second);

  fmt.Printf("Got here: %d\n", 4);
  device.SetLEDColor(button,"blue")
  time.Sleep(time.Second);

  for i := 0; i < int(howler.ButtonMax); i++ {
    c, _ := device.GetLEDColor(howler.Buttons(i))
    fmt.Printf("%03d: %s\n", i, c.ToHexString())
  }

  fw, _ := device.GetFWRelease()
  fmt.Printf("Firmware: %d.%d\n", fw.Major, fw.Minor)

  for i := 0; i < 200; i++ {
    data, err := device.ReadAccel()
    if err != nil {
      fmt.Println(err.Error())
    } else {
      fmt.Println(data.String())
    }
  }

 device.SetInput(howler.InputButton26, howler.ModeKeyboard, howler.KeyZ, howler.ModifierNone)
 input, err := device.GetInput(howler.InputButton26)
 input.Dump()

  for i := 0; i < int(howler.InputMax); i++ {
   _, err := device.GetInput(howler.Inputs(i))
    if err != nil {
      fmt.Println(err.Error())
    }
  }
*/
