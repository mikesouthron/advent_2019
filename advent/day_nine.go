package advent

import (
	"strconv"
	"strings"
)

func RunDayNine() {
	data := strings.Split(ReadFile(9)[0], ",")
	dataValues := make([]int, 4294967295)
	for i, input := range data {
		val, _ := strconv.Atoi(strings.TrimSpace(input))
		dataValues[i] = val
	}
	x, y, _ := RunIntCode(dataValues, 2, 1, 0)
	println(x)
	println(y)
}
