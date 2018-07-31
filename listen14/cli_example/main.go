package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// func main() {
// 	app := cli.NewApp()
// 	app.Name = "hello"
// 	app.Usage = "hello world"
// 	app.Action = func(c *cli.Context) error {
// 		fmt.Println("Hello my friend")
// 		fmt.Printf("hello %q", c.Args().Get(0))
// 		return nil
// 	}
// 	app.Run(os.Args)
// }

// func main() {
// 	app := cli.NewApp()
// 	app.Name = "hello"
// 	app.Usage = "hello world"
// 	app.Action = func(c *cli.Context) error {
// 		var cmd string
// 		if c.NArg() > 0 {
// 			cmd = c.Args()[0]
// 		}
// 		fmt.Println("hello frine cmd:", cmd)
// 		return nil
// 	}
// 	app.Run(os.Args)
// }

func main() {
	app := cli.NewApp()
	app.Name = "Cli - test"
	app.Usage = "hello CLI"
	var (
		language string
		flag     bool
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang,l",
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "boool,b",
			Usage:       "select boolean",
			Destination: &flag,
		},
	}
	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Printf("cmd is :", cmd)
		}
		fmt.Println("language is", language)
		fmt.Println("flag is :", flag)
		return nil
	}
	app.Run(os.Args)
}
