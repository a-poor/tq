package main

import (
  "fmt"
  "os"
)

func main() {
  app := makeApp()
  if err := app.Run(os.Args); err != nil {
    fmt.Fprintf(os.Stderr, "error: %v\n", err)
    os.Exit(1)
  }
}

