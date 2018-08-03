package generateGovBall

import (
	"github.com/study/twoColorBall/generateBall"
)

func GenerateGovBall() ([]int, []int) {
	baseRed, baseBlue := generateBall.GenerateBall(0), generateBall.GenerateBall(1)
	return baseRed, baseBlue
}
