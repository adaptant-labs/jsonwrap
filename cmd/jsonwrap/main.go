package main

import (
	"fmt"
	"github.com/adaptant-labs/jsonwrap"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	var destinationUrl string
	var method string
	var wrappers cli.StringSlice

	app := &cli.App{
		Name:        "jsonwrap",
		Usage:       "CLI for JSON object nesting and HTTP forwarding",
		ArgsUsage:   "<JSON file>",
		Description: "A convenience CLI tool for nesting JSON objects and forwarding the result to a remote HTTP endpoint",
		Author:      "Adaptant Labs <labs@adaptant.io>",
		Copyright:   "(c) 2019 Adaptant Solutions AG",
		Version:     "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "target",
				Usage:       "Target URL to forward wrapped JSON to",
				Destination: &destinationUrl,
			},
			&cli.StringFlag{
				Name:        "method",
				Usage:       "Forwarding method to use",
				Value:       "POST",
				Destination: &method,
			},
			&cli.StringSliceFlag{
				Name:     "wrap",
				Usage:    "Name of object to nest under",
				Value:    &wrappers,
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			var jsonStr string
			var err error

			if len(c.Args()) == 0 {
				// Read from stdin
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					jsonStr, err = jsonFromStdin()
					if err != nil {
						return err
					}
				} else {
					return cli.ShowAppHelp(c)
				}
			} else {
				// Read from file
				jsonStr, err = jsonFromFile(c.Args().Get(0))
				if err != nil {
					return err
				}
			}

			j := jsonwrap.NewJSONWrapper()

			// Apply all wrap values
			for _, w := range wrappers.Value() {
				jsonStr = j.Wrap(w, jsonStr)
			}

			if destinationUrl != "" {
				if !strings.EqualFold(method, "POST") && !strings.EqualFold(method, "PUT") {
					return fmt.Errorf("invalid forwarding method defined - must be one of POST or PUT")
				}

				return jsonForwardToUrl(jsonStr, strings.ToUpper(method), destinationUrl)
			}

			fmt.Println(jsonStr)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
