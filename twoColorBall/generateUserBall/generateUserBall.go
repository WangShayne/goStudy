package generateUserBall

import (
	"github.com/study/twoColorBall/generateBall"
)

type Ball struct {
	RedBalls  []int
	BlueBalls []int
	RedWin    int
	BlueWin   int
	Bonus     int
}

type userInfo struct {
	Balls []Ball
	Total int
	Pay   int
}

func GenerateUserBall(length int) userInfo {
	if length == 0 {
		length = 1
	}

	// UserRed, UserBlue := generateBall.GenerateBall(0), generateBall.GenerateBall(1)
	var user userInfo
	for i := 0; i < length; i++ {
		var userBall Ball
		userBall.RedBalls = append(userBall.RedBalls, generateBall.GenerateBall(0)...)
		userBall.BlueBalls = append(userBall.BlueBalls, generateBall.GenerateBall(1)...)

		user.Balls = append(user.Balls, userBall)
	}
	return user
}
