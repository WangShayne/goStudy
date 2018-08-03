package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/study/twoColorBall/generateGovBall"
	"github.com/study/twoColorBall/generateUserBall"
)

func main() {
	app := cli.NewApp()
	app.Name = "dobule color ball"
	app.Usage = "test dobule color ball"

	var length int

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "zhu,z",
			Usage:       "input number",
			Destination: &length,
		},
	}
	app.Action = func(c *cli.Context) error {
		govRed, govBlue := generateGovBall.GenerateGovBall()
		user := generateUserBall.GenerateUserBall(length)
		user.Pay = 2 * length
		// for

		fmt.Printf("%v", user)
		fmt.Println(user.Pay)
		fmt.Println(govRed, govBlue)

		return nil
	}
	app.Run(os.Args)

}
