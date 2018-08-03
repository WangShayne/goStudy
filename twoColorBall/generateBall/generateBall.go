package generateBall

import (
	"math/rand"
	"sort"
	"time"
)

func GenerateBall(tp int) []int {
	var (
		redBallList  []int
		blueBallList []int
		redBalls     []int
		blueBalls    []int
	)
	for i := 1; i < 33; i++ {
		if i < 17 {
			blueBallList = append(blueBallList, i)
		}
		redBallList = append(redBallList, i)
	}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 6; i++ {
		n := rand.Intn(len(redBallList))
		redBalls = append(redBalls, redBallList[n])
	}
	x := rand.Intn(len(blueBallList))
	blueBalls = append(blueBalls, blueBallList[x])
	sort.Ints(redBalls[:])
	sort.Ints(blueBalls[:])
	if tp == 0 {
		return redBalls
	} else {
		return blueBalls
	}
}
