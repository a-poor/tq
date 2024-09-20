package main

import (
	"fmt"
	"slices"

	"github.com/urfave/cli/v2"
)

func makeApp() *cli.App {
  return &cli.App{
    Name: "tq",
    Usage: "",
    Version: "0.0.1",
    Flags: []cli.Flag{
      &cli.StringFlag{
        Name: "input-file",
        Aliases: []string{"i"},
        Usage: "input file (\"-\" for stdin)",
      },
      &cli.StringFlag{
        Name: "input-type",
        Usage: "input file type (infer from file extension if not specified)",
        Action: func(c *cli.Context, v string) error {
          validTypes := []string{"json", "csv", "xlsx"}
          if !slices.Contains(validTypes, v) {
            return fmt.Errorf("invalid input type: %s", v)
          }
          return nil
        },
      },
      &cli.StringFlag{
        Name: "input-sheet",
        Usage: "input sheet name (if applicable)",
      },
      &cli.StringFlag{
        Name: "output-file",
        Aliases: []string{"o"},
        Usage: "output file (default: stdout)",
      },
      &cli.StringFlag{
        Name: "output-type",
        Usage: "output file type (infer from file extension if not specified)",
        Action: func(c *cli.Context, v string) error {
          validTypes := []string{"json", "csv", "xlsx"}
          if !slices.Contains(validTypes, v) {
            return fmt.Errorf("invalid output type: %s", v)
          }
          return nil
        },
      },
      &cli.StringFlag{
        Name: "output-sheet",
        Usage: "output sheet name (if applicable)",
      },
      &cli.StringFlag{
        Name: "columns",
        Usage: "columns to include in output (comma-separated)",
      },
    },
    Action: func(c *cli.Context) error {
      return nil
    },
  }
}
