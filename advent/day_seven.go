package advent

import (
	"fmt"
	"strconv"
	"strings"
)

//ExecuteDaySeven run day seven
func ExecuteDaySeven() {
	data := strings.Split(ReadFile(7)[0], ",")
	var dataValues []int
	for _, input := range data {
		val, _ := strconv.Atoi(strings.TrimSpace(input))
		dataValues = append(dataValues, val)
	}

	inputs := make([]int, 5)
	inputs[0] = 5
	inputs[1] = 6
	inputs[2] = 7
	inputs[3] = 8
	inputs[4] = 9

	permuts := permutation(inputs)

	maxOutput := 0

	for _, current := range permuts {
		output := getOutput(current, dataValues)
		if output > maxOutput {
			maxOutput = output
		}
	}

	fmt.Println(maxOutput)
}

func permutation(xs []int) (permuts [][]int) {
	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

type stateS struct {
	data []int
	pos  int
}

func getOutput(inputs []int, dataValues []int) int {
	input := 0
	state := make([]stateS, 5)

	for i, value := range inputs {
		data := make([]int, len(dataValues))
		copy(data, dataValues)
		newpos := 0
		input, newpos, _ = RunIntCode(data, value, input, 0)
		state[i] = stateS{data, newpos}
	}

	finished := false

	for {
		for i := 0; i < 5; i++ {
			newpos := 0
			data := state[i].data
			previousInput := input
			input, newpos, finished = RunIntCode(data, input, input, state[i].pos)
			state[i] = stateS{data, newpos}
			if i == 4 && finished {
				return previousInput
			}
			if input == 0 {
				input = previousInput
			}
		}
	}
}
