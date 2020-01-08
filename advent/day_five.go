package advent

import (
	"math"
	"strconv"
	"strings"
)

func ExecuteDayFive() {
	data := strings.Split(ReadFile(5)[1], ",")
	var dataValues []int
	for _, input := range data {
		val, _ := strconv.Atoi(strings.TrimSpace(input))
		dataValues = append(dataValues, val)
	}
	RunIntCode(dataValues, 5, 5, 0)
}

//func execute(data []int) (int, int) {
//	for n := 0; n <= 99; n++ {
//		for v := 0; v <= 99; v++ {
//			runValues := make([]int, len(data))
//			copy(runValues, data)
//			runValues[1] = n
//			runValues[2] = v
//			output := run(runValues)
//			if output == 19690720 {
//				return n, v
//			}
//		}
//	}
//	return 0, 0
//}

func RunIntCode(data []int, input, input_two, pos int) (int, int, bool) {
	one := true
	output := 0
	finished := false
	for {
		s := strconv.Itoa(data[pos])
		l := len(s)
		opCode, _ := strconv.Atoi(s[int(math.Max(float64(l-2), 0)):])

		if opCode == 99 {
			return 0, 0, true
		}

		modeOne := 0
		modeTwo := 0
		modeThree := 0

		if l - 3 >= 0 {
			modeOne, _ = strconv.Atoi(s[l-3 : l-2])
			if l - 4 >= 0 {
				modeTwo, _ = strconv.Atoi(s[l-4 : l-3])
				if l - 5 >= 0 {
					modeThree, _ = strconv.Atoi(s[l-5 : l-4])
				}
			}
		}

		if opCode == 1 || opCode == 2 {
			if modeThree == 1 {
				panic("Should not be possible!")
			}
			posOne := data[pos+1]
			var inputOne int
			if modeOne == 0 {
				inputOne = data[posOne]
			} else {
				inputOne = posOne
			}

			posTwo := data[pos+2]
			var inputTwo int
			if modeTwo == 0 {
				inputTwo = data[posTwo]
			} else {
				inputTwo = posTwo
			}

			updatePos := data[pos+3]
			if opCode == 1 {
				data[updatePos] = inputOne + inputTwo
			} else if opCode == 2 {
				data[updatePos] = inputOne * inputTwo
			}

			pos = pos + 4
		}

		if opCode == 3 {
			updatePos := data[pos+1]
			if one {
				data[updatePos] = input
				one = false
			} else {
				data[updatePos] = input_two
			}

			pos = pos + 2
		}

		if opCode == 4 {
			if modeOne == 0 {
				return data[data[pos+1]], pos + 2, false
			} else {
				return data[pos+1], pos + 2, false
			}
		}

		if opCode == 5 {
			one := getValue(data, pos, 1, modeOne)
			two := getValue(data, pos, 2, modeTwo)
			if one != 0 {
				pos = two
			} else {
				pos = pos + 3
			}
		}

		if opCode == 6 {
			one := getValue(data, pos, 1, modeOne)
			two := getValue(data, pos, 2, modeTwo)
			if one == 0 {
				pos = two
			} else {
				pos = pos + 3
			}
		}

		if opCode == 7 {
			one := getValue(data, pos, 1, modeOne)
			two := getValue(data, pos, 2, modeTwo)
			three := data[pos+3]
			//three := getValue(data, pos, 3, modeThree)
			if one < two {
				data[three] = 1
			} else {
				data[three] = 0
			}
			pos = pos + 4
		}

		if opCode == 8 {
			one := getValue(data, pos, 1, modeOne)
			two := getValue(data, pos, 2, modeTwo)
			three := data[pos+3]
			//three := getValue(data, pos, 3, modeThree)
			if one == two {
				data[three] = 1
			} else {
				data[three] = 0
			}
			pos = pos + 4
		}
	}

	return output, pos, finished
}

func getValue(data []int, pos, index, mode int) int {
	if mode == 0 {
		return data[data[pos+index]]
	} else {
		return data[pos+index]
	}
}