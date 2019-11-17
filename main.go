package main

import (
  "os"
  "log"
  "fmt"

  _"sort"
  "io/ioutil"

  "gopkg.in/urfave/cli.v2"
  howler "github.com/sferris/howler-controller"
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
      // \
      //  > Set inputs LED color
      // /
      {
        Name:        "set-led",
        Usage:       "Change the color of one of the Button LEDs",
        Description: "This command is used to alter the color of a button/LED",

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
            Usage: "The scope to change the color (current, or default)",
          },
        },

        Action: func(c *cli.Context) error {
          led := LEDStruct{
            Name:   c.String("button"),
            Color:  c.String("color"),
            Scope:  c.String("scope"),
          }
          return led.Process()
        },
      },

      // \
      //  > Set input to emit a Keyboard value
      // /
      {
        Name:        "set-input-keyboard",
        Aliases:     []string{"set-keyboard"},
        Usage:       "Alter a controller input to emit a keyboard key",
        Description: "This command will set a controller input to emit a keyboard key",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "modifier",
            Usage: "In a keyboard context, the modifier to use in addition to the value",
          },
          &cli.StringFlag{
            Name: "key",
            Usage: "What value the button emits when pressed (context dependent)",
          },
        },

        Action: func(c *cli.Context) error {
          input := InputStruct{
            Name:      c.String("input"),
            Type:      "keyboard",
            Modifier:  c.String("modifier"),
            Value:     c.String("key"),
          }
          return input.Process()
        },
      },

      // \
      //  > Set input to emit a Joystick button
      // /
      {
        Name:        "set-input-joystick",
        Aliases:     []string{"set-joystick"},
        Usage:       "Alter a controller input to emit a joystick button",
        Description: "This command will set a controller input to emit a joystick button press",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "joystick",
            Value: "joystick1",
            Usage: "The joystick number to emit, eg: 1 or 2",
          },
          &cli.StringFlag{
            Name: "button",
            Usage: "Which joystick button to emit",
          },
        },

        Action: func(c *cli.Context) error {
          input := InputStruct{
            Name:      c.String("input"),
            Type:      c.String("joystick"),
            Value:     c.String("button"),
          }
          return input.Process()
        },
      },

      // \
      //  > Set input to emit a Mouse button
      // /
      {
        Name:        "set-input-mouse",
        Aliases:     []string{"set-mouse"},
        Usage:       "Alter a controller input to emit a mouse button",
        Description: "This command will set a controller input to emit a mouse button press",

        Flags: []cli.Flag{
          &cli.StringFlag{
            Name: "input",
            Usage: "The input to alter, eg: joy1up, joy1down, button1, etc",
          },
          &cli.StringFlag{
            Name: "button",
            Usage: "Which mouse button to emit",
          },
        },

        Action: func(c *cli.Context) error {
          input := InputStruct{
            Name:      c.String("input"),
            Type:      "keyboard",
            Modifier:  c.String("modifier"),
            Value:     c.String("button"),
          }
          return input.Process()
        },
      },

      // \
      //  > Process many inputs defined in a file
      // /
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
          file := FileStruct{
            Path: c.String("path"),
          }
          return file.Process()
        },
      },

      // \
      //  > Show possible keyboard key values
      // /
      {
        Name:        "show-keyboard-keys",
        Usage:       "Show the valid keyboard keys",
        Aliases:     []string{"show-keys"},
        Description: "This shows the list of valid keyboard key names",

        Action: func(c *cli.Context) error {
          fmt.Println("Valid keyboard keys:")
          //fmt.Println(strings.Join(howler.KeyNames[howler.KeyMin:], ", "))
          return nil
        },
      },

      // \
      //  > Show possible keyboard modifier values
      // /
      {
        Name:        "show-keyboard-modifiers",
        Usage:       "Show the valid keyboard modifier names",
        Aliases:     []string{"show-modifiers"},
        Description: "This shows the list of valid keyboard modifier names",

        Action: func(c *cli.Context) error {
          fmt.Println("Valid keyboard modifier names:")
          //fmt.Println(strings.Join(howler.ModifierNames[howler.ModifierMin:], ", "))
          return nil
        },
      },

      // \
      //  > Show possible mouse button values
      // /
      {
        Name:        "show-mouse-buttons",
        Usage:       "Show the valid mouse buttons",
        Aliases:     []string{"show-mouse"},
        Description: "This shows the list of valid mouse buttons",

        Action: func(c *cli.Context) error {
          fmt.Println("Valid mouse buttons:")
          //fmt.Println(strings.Join(howler.MouseNames[howler.MouseMin:], ", "))
          return nil
        },
      },

      // \
      //  > Show possible joystick button values
      // /
      {
        Name:        "show-joystick-buttons",
        Usage:       "Show the valid joystick buttons",
        Aliases:     []string{"show-joystick"},
        Description: "This shows the list of valid joystick buttons",

        Action: func(c *cli.Context) error {
          fmt.Println("Valid joystick buttons:")
          //fmt.Println(strings.Join(howler.JoystickButtonNames[howler.JoyMin:], ", "))
          return nil
        },
      },

      // \
      //  > Show firmware version
      // /
      {
        Name:        "firmware",
        Usage:       "Show the firmware revision information",
        Aliases:     []string{"show-firmward"},
        Description: "This shows the list of valid joystick buttons",

        Action: func(c *cli.Context) error {
          return ShowFirmware()
        },
      },
    },
}

var controller *howler.HowlerDevice

func main() {
  log.SetFlags(0)
  log.SetOutput(ioutil.Discard)

/*
  var keys []int
  for k := range howler.ModifierNames {
      keys = append(keys, int(k))
  }
  sort.Ints(keys)

  // To perform the opertion you want
  for _, k := range keys {
      fmt.Printf("Key: %0x, Value: %s\n", k, howler.ModifierNames[howler.Modifiers(k)])
  }
*/

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err.Error())
  }

  if controller != nil {
    log.Printf("Closing howler device")
    controller.Close()
  }
}

