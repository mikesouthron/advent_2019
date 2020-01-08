package advent

import (
	"strconv"
	"strings"
)

func DayTwoPartOneExecute() {
	data := strings.Split(ReadFile(2)[0], ",")
	var dataValues []int
	for _, input := range data {
		val, _ := strconv.Atoi(strings.TrimSpace(input))
		dataValues = append(dataValues, val)
	}
	noun, verb := execute(dataValues)
	println(100 * noun + verb)
}

func execute(data []int) (int, int) {
	for n := 0; n <= 99; n++ {
		for v := 0; v <= 99; v++ {
			runValues := make([]int, len(data))
			copy(runValues, data)
			runValues[1] = n
			runValues[2] = v
			output := run(runValues)
			if output == 19690720 {
				return n, v
			}
		}
	}
	return 0, 0
}

func run(data []int) int {
	pos := 0
	for {
		opCode := data[pos]
		if opCode == 99 {
			break
		}

		posOne := data[pos+1]
		inputOne := data[posOne]
		posTwo := data[pos+2]
		inputTwo := data[posTwo]
		updatePos := data[pos+3]

		if opCode == 1 {
			data[updatePos] = inputOne + inputTwo
		} else if opCode == 2 {
			data[updatePos] = inputOne * inputTwo
		}

		pos = pos + 4
	}

	return data[0]
}
