package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/study/twoColorBall/generateGovBall"
	"github.com/study/twoColorBall/generateUserBall"
)

func checkWin(a []int, b []int) int {
	var win int = 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				win++
				break
			}
		}
	}
	return win
}

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

		if length == 0 {
			length = 1
		}
		fmt.Printf("您购买了:%d注双色球彩票\n", length)
		user := generateUserBall.GenerateUserBall(length)
		user.Pay = 2 * length
		fmt.Printf("您支付了:%d元\n\n", user.Pay)

		govRed, govBlue := generateGovBall.GenerateGovBall()
		fmt.Printf("本期开奖号码为:%s %s\n\n", strings.Replace(strings.Trim(fmt.Sprint(govRed), "[]"), " ", ",", -1), strings.Replace(strings.Trim(fmt.Sprint(govBlue), "[]"), " ", ",", -1))
		for i := 0; i < len(user.Balls); i++ {

			fmt.Printf("您的投注号码为:%s %s\n", strings.Replace(strings.Trim(fmt.Sprint(user.Balls[i].RedBalls), "[]"), " ", ",", -1), strings.Replace(strings.Trim(fmt.Sprint(user.Balls[i].BlueBalls), "[]"), " ", ",", -1))
			user.Balls[i].RedWin = checkWin(user.Balls[i].RedBalls, govRed)
			user.Balls[i].BlueWin = checkWin(user.Balls[i].BlueBalls, govBlue)
			if user.Balls[i].BlueWin == 1 {
				switch user.Balls[i].RedWin {
				case 1, 0:
					user.Total += 5
				case 2:
					user.Total += 10
				case 3:
					user.Total += 20
				case 4:
					user.Total += 200
				case 5:
					user.Total += 20000
				case 6:
					user.Total += 5000000
				}
			} else {
				switch user.Balls[i].RedWin {
				case 2, 1, 0:
					user.Total += 0
				case 3:
					user.Total += 5
				case 4:
					user.Total += 20
				case 5:
					user.Total += 200
				case 6:
					user.Total += 2000
				}
			}
		}
		fmt.Printf("\n您的中奖金额为:%d\n", user.Total)
		return nil
	}
	app.Run(os.Args)

}
