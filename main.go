package main

import (
  "os"
  "fmt"
  "gopkg.in/urfave/cli.v2"
)

var (
  device int
)

var app = &cli.App{
    Name: "howler-cli",
    Usage: "A command line utility for configuring a Howler Arcade controller",
    Version: "0.0.1",

    Flags: []cli.Flag{
      &cli.IntFlag{
        Name: "device",
        Aliases: []string{"d"},
        Value: 0,
        Usage: "The Howler Controller to configure (0..n)",
        Destination: &device,
      },
    },

    Commands: []*cli.Command{
      {
        Name:        "set-led",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This comment is used to alter the color of a button/LED",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "button",
            Usage: "The button/led to change",
          },
          &cli.StringFlag{
            Name: "color",
            Usage: "The color to set the button/led to",
          },
          &cli.StringFlag{
            Name: "scope",
            Value: "current",
            Usage: "The scope to change the color (current/default)",
          },
        },

        Action: func(c *cli.Context) error {
          fmt.Printf("Device: %d\n", device)
          fmt.Printf(
            "Button: %s, Color: %s, Scope: %s\n\n", 
              c.String("button"),
              c.String("color"),
              c.String("scope"),
          )
          return nil
        },
      },

      {
        Name:        "set-input",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This command is used to alter the behavior of a button on the controller",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "button",
            Usage: "The button/led to change, eg: button1",
          },
          &cli.StringFlag{
            Name: "mode",
            Usage: "The context used by the button when emitting a value (joystick1 or 2, keyboard, mouse)",
          },
          &cli.StringFlag{
            Name: "modifier",
            Value: "none",
            Usage: "In keyboard mode, the modifier to use in addition to the value. ([left|right] control, shift, alt, ui)",
          },
          &cli.StringFlag{
            Name: "value",
            Usage: "What value the button emits when pressed (context dependent)",
          },
        },

        Action: func(c *cli.Context) error {
          fmt.Printf("Device: %d\n", device)
          fmt.Printf(
            "Button: %s, Mode: %s, Modifier: %s, Value: %s\n\n", 
              c.String("button"),
              c.String("mode"),
              c.String("modifier"),
              c.String("value"),
          )
          return nil
        },
      },

      {
        Name:        "from-file",
        Usage:       "Read settings from a yaml file",
        Aliases:     []string{"read"},
        Description: "This command allows you to change many settings that are represented in a yaml file",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "path",
            Aliases: []string{"file"},
            Usage: "The fullpath to the yaml file containing the settings to be applied",
          },
        },

        Action: func(c *cli.Context) error {
          fmt.Printf("Device: %d\n", device)
          fmt.Printf(
            "Filename: %s\n\n",
              c.String("path"),
          )
          return nil
        },

      },
    },
}

func main() {
  app.Run(os.Args)
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
