package main

import (
  "fmt"
  "regexp"
  "strconv"
  "github.com/sferris/howler-controller/color"
)

type RGBStruct struct {
  Red, Green, Blue int
}

func (rgb *RGBStruct) String() string {
  return fmt.Sprintf("Red: %03d, Green: %03d, Blue: %03d", rgb.Red, rgb.Green, rgb.Blue)
}

var colors map[string]RGBStruct

func FetchRGB(value string) (RGBStruct, bool) {
  re := regexp.MustCompile("(?:rgb\\s*:\\s*)(\\d+),(\\d+),(\\d+)")

  match := re.FindStringSubmatch(value)
  if match != nil {
    r, _ := strconv.Atoi(match[1])
    g, _ := strconv.Atoi(match[2])
    b, _ := strconv.Atoi(match[3])

    return RGBStruct{Red: r, Green: g, Blue: b}, true
  }

  if result, ok := colors[value]; ok {
    return result, true
  }

  if rgb, ok := color.Lookup(value); ok {
    return RGBStruct{Red: rgb.Red, Green: rgb.Green, Blue: rgb.Blue}, true;
  }

  return RGBStruct{}, false;
}
