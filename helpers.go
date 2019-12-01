package main

import (
  "fmt"

  "sort"

  howler "github.com/sferris/howler-controller"
)

func ControlInputNames() string {
  var result string

  w := 0
  for _, control := range howler.ControlInputsByID() {
    value := fmt.Sprintf("%s, ", control.Name())

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintf("\n")
    }
  }

  return result
}

func ControlInputNameByCapability(capability howler.ControlCapability) string {
  var result string

  w := 0
  for _, control := range howler.ControlInputsByID() {
    if control.Capability() & capability != 0 {
      value := fmt.Sprintf("%s, ", control.Name())

      result += value

      w += len(value)
      if w >= (columns) {
        w=0
        result += fmt.Sprintf("\n")
      }
    }
  }

  return result
}

func ControlFunctionNames() string {
  var result string

  w := 0
  for _, control := range howler.ControlFunctionsByID() {
    value := fmt.Sprintf("%s, ", control.Name())

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintf("%s")
    }
  }

  return result
}

func ControlFunctionNameByCapability(capability howler.ControlCapability) string {
  var result string

  w := 0
  for _, control := range howler.ControlFunctionsByID() {
    if control.Capability() & capability != 0 {
      value := fmt.Sprintf("%s, ", control.Name())

      result += value

      w += len(value)
      if w >= (columns) {
        w=0
        result += fmt.Sprintf("\n")
      }
    }
  }

  return result
}

func JoystickButtonNames() string {
  var result string

  w := 0
  for _, name := range howler.JoystickButtonNames {
    value := fmt.Sprintf("%s, ", name)

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintf("\n")
    }
  }

  return result
}

func KeyNames() string {
  var result string

  var keys []int
  for k := range howler.KeyNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.KeyNames[howler.KeyCodes(k)])

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result
}
func ModifierNames() string {
  var result string

  var keys []int
  for k := range howler.ModifierNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.ModifierNames[howler.KeyModifiers(k)])

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result
}

func MouseButtonNames() string {
  var result string

  var keys []int
  for k := range howler.MouseButtonNames {
    keys = append(keys, int(k))
  }

  sort.Ints(keys)

  w := 0
  for _, k := range keys {
    value := fmt.Sprintf("%s, ", howler.MouseButtonNames[howler.MouseButtons(k)])

    result += value

    w += len(value)
    if w >= (columns) {
      w=0
      result += fmt.Sprintln()
    }

  }

  return result
}
