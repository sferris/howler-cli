package main

import (
  "os"
  "fmt"
  "log"
  "flag"
  "strings"
  "strconv"

  _"time"

  howler "github.com/sferris/howler-controller"
)

type rgbFlags struct {
  Red   int
  Blue  int
  Green int
}

func (rgb *rgbFlags) String() string {
  return "foo"
}

func (rgb *rgbFlags) Set(value string) error {
  parsed := strings.Split(value, ",")

  var red, blue, green int
  var err error

  if red, err = strconv.Atoi(parsed[0]); err != nil {
    log.Fatalf("Invalid red value: %s\n", parsed[0])
  }
  if green, err = strconv.Atoi(parsed[1]); err != nil {
    log.Fatalf("Invalid green value: %s\n", parsed[1])
  }
  if blue, err = strconv.Atoi(parsed[2]); err != nil {
    log.Fatalf("Invalid blue value: %s\n", parsed[2])
  }
  *rgb = rgbFlags{red, green, blue}
  return nil
}

type Input struct {
  Button    string
  Mode      string
  Modifier  string
  Value     string
}

func (input *Input) String() string {
  switch strings.ToLower(input.Mode) {
    case "joystick1": fallthrough
    case "joystick2":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          input.Value,
      )
    case "keyboard":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          howler.Key(input.Value),
      )
    case "mouse":
      return fmt.Sprintf(
        "Button: %d, Mode: %d, Modifier: %d, Value: %d",
          howler.Input(input.Button),
          howler.Mode(input.Mode),
          howler.Modifier(input.Modifier),
          howler.MouseButton(input.Value),
      )
  }
  return ""
}

func (input *Input) Set() error {
  fmt.Println(input)
  return nil
}

type Led struct {
  Button    string
  Scope     string
  RGB       rgbFlags
}

func (led *Led) String() string {
  return fmt.Sprintf(
    "Button: %d, Scope: %s, Red: %d, Green: %d, Blue: %d",
      howler.Button(led.Button),
      led.Scope,
      led.RGB.Red,
      led.RGB.Green,
      led.RGB.Blue,
  )
}

func (led *Led) Set() error {
  fmt.Println(led)
  return nil
}

type Work struct {
  Game    string
  Leds    []Led
  Inputs  []Input
}

func main() {
  setLedCMD := flag.NewFlagSet("set-led", flag.ExitOnError)
  setInputCMD := flag.NewFlagSet("set-input", flag.ExitOnError)
  setReadFileCMD := flag.NewFlagSet("read-file", flag.ExitOnError)

  LedButton := setLedCMD.String("led", "", "Led Button to change the color (Required)")
  LedScope   := setLedCMD.String("scope", "current", "The Led scope (default or current)")

  var LedRGB rgbFlags
  setLedCMD.Var(&LedRGB, "rgb", "The RGB value for the Led color (Required)")

  InputButton   := setInputCMD.String("input", "", "Button to set the Input properties (Required)")
  InputMode     := setInputCMD.String("mode", "", "The Input mode [Joystick1 or 2||Keyboard||Mouse] (Required)")
  InputModifier := setInputCMD.String("modifier", "", "The Input key modifier")
  InputValue    := setInputCMD.String("value", "", "The Input value (Required)")

  FileName := setReadFileCMD.String("file", "",
    "File to read instructions from (Required)")

  if len(os.Args) < 2 {
    flag.PrintDefaults()
    os.Exit(1)
  }

  switch os.Args[1] {
    case "set-led":
      setLedCMD.Parse(os.Args[2:])

    case "set-input":
      setInputCMD.Parse(os.Args[2:])

    case "read-file": 
      setReadFileCMD.Parse(os.Args[2:])

    default:
      flag.PrintDefaults()
      os.Exit(1)
  }

  if setLedCMD.Parsed() {
    led := Led{
      Button: *LedButton,
      Scope:  *LedScope,
      RGB:     LedRGB,
    }

    led.Set();
  }

  if setInputCMD.Parsed() {
    input := Input{
      Button:    *InputButton,
      Mode:      *InputMode,
      Modifier:  *InputModifier,
      Value:     *InputValue,
    }

    input.Set();
  }

  if setReadFileCMD.Parsed() {
    fmt.Printf("Filename: %s\n", *FileName)
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
