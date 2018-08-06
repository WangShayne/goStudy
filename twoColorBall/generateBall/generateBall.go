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

	for i := 0; i < 100; i++ {
		n := rand.Intn(len(redBallList))
		if len(redBalls) == 0 {
			redBalls = append(redBalls, redBallList[n])
		} else if len(redBalls) == 6 {
			break
		} else {
			in := false
			for j := 0; j < len(redBalls); j++ {
				if redBallList[n] == redBalls[j] {
					in = true
				} else {
					break
				}
			}
			if !in {
				redBalls = append(redBalls, redBallList[n])
			}
		}
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
