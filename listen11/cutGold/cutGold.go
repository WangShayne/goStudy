package cutGold

import (
	"strings"
)

func Cut(total int, users []string, distributeion map[string]int) (result int, resultMap map[string]int) {

	for _, value := range users {
		_, ok := distributeion[value]
		if !ok {
			distributeion[value] = 0
		}
	}

	for index, value := range distributeion {
		if strings.Index(index, "a") > 0 || strings.Index(index, "A") > 0 || strings.Index(index, "e") > 0 || strings.Index(index, "E") > 0 {
			value++
			total--
		} else if strings.Index(index, "i") > 0 || strings.Index(index, "I") > 0 {
			value += 2
			total -= 2
		} else if strings.Index(index, "o") > 0 || strings.Index(index, "O") > 0 {
			value += 3
			total -= 3
		} else if strings.Index(index, "u") > 0 || strings.Index(index, "U") > 0 {
			value += 5
			total -= 5
		}
		distributeion[index] = value
	}
	// fmt.Println(distributeion)
	result = total
	resultMap = distributeion

	return result, resultMap
}
