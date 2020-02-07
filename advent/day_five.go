package advent

import (
	"strconv"
	"strings"
)

//ExecuteDayFive run day 5
func ExecuteDayFive() {
	data := strings.Split(ReadFile(5)[1], ",")
	var dataValues []int
	for _, input := range data {
		val, _ := strconv.Atoi(strings.TrimSpace(input))
		dataValues = append(dataValues, val)
	}
	RunIntCode(dataValues, 5, 5, 0)
}
