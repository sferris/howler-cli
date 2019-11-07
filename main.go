package main

import (
  "os"
  _"fmt"
  _"log"
  "flag"
  _"strings"
  _"strconv"

  _"time"

  //howler "github.com/sferris/howler-controller"
)

func main() {
  if len(os.Args) < 2 {
    flag.PrintDefaults()
    os.Exit(1)
  }

  switch os.Args[1] {
    case "set-led":
      parseLed()

    case "set-input":
      parseInput()

    case "read-file": 
      parseYAML()

    default:
      flag.PrintDefaults()
      os.Exit(1)
  }

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
