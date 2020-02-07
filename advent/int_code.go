package advent

import (
	"math"
	"strconv"
)

//RunIntCode run int_code machine
func RunIntCode(data []int, input, inputTwo, pos int) (int, int, bool) {
	one := true
	relativeBase := 0
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

		if l-3 >= 0 {
			modeOne, _ = strconv.Atoi(s[l-3 : l-2])
			if l-4 >= 0 {
				modeTwo, _ = strconv.Atoi(s[l-4 : l-3])
				if l-5 >= 0 {
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
			switch modeOne {
			case 0:
				inputOne = data[posOne]
			case 1:
				inputOne = posOne
			case 2:
				inputOne = data[posOne+relativeBase]
			}

			posTwo := data[pos+2]
			var inputTwo int

			switch modeTwo {
			case 0:
				inputTwo = data[posTwo]
			case 1:
				inputTwo = posTwo
			case 2:
				inputTwo = data[posTwo+relativeBase]
			}

			updatePos := data[pos+3]
			if opCode == 1 {
				if modeThree == 2 {
					data[updatePos+relativeBase] = inputOne + inputTwo
				} else {
					data[updatePos] = inputOne + inputTwo
				}
			} else if opCode == 2 {
				if modeThree == 2 {
					data[updatePos+relativeBase] = inputOne * inputTwo
				} else {
					data[updatePos] = inputOne * inputTwo
				}
			}

			pos = pos + 4
		}

		if opCode == 3 {
			updatePos := data[pos+1]
			if one {
				if modeOne == 2 {
					data[updatePos+relativeBase] = input
				} else {
					data[updatePos] = input
				}
				one = false
			} else {
				if modeOne == 2 {
					data[updatePos+relativeBase] = inputTwo
				} else {
					data[updatePos] = inputTwo
				}
			}

			pos = pos + 2
		}

		if opCode == 4 {
			switch modeOne {
			case 0:
				return data[data[pos+1]], pos + 2, false
			case 1:
				return data[pos+1], pos + 2, false
			case 2:
				return data[data[pos+1]+relativeBase], pos + 2, false
			}
		}

		if opCode == 5 {
			one := getValue(data, pos, 1, modeOne, relativeBase)
			two := getValue(data, pos, 2, modeTwo, relativeBase)
			if one != 0 {
				pos = two
			} else {
				pos = pos + 3
			}
		}

		if opCode == 6 {
			one := getValue(data, pos, 1, modeOne, relativeBase)
			two := getValue(data, pos, 2, modeTwo, relativeBase)
			if one == 0 {
				pos = two
			} else {
				pos = pos + 3
			}
		}

		if opCode == 7 {
			one := getValue(data, pos, 1, modeOne, relativeBase)
			two := getValue(data, pos, 2, modeTwo, relativeBase)
			three := data[pos+3]
			if one < two {
				if modeThree == 2 {
					data[three+relativeBase] = 1
				} else {
					data[three] = 1
				}
			} else {
				if modeThree == 2 {
					data[three+relativeBase] = 0
				} else {
					data[three] = 0
				}
			}
			pos = pos + 4
		}

		if opCode == 8 {
			one := getValue(data, pos, 1, modeOne, relativeBase)
			two := getValue(data, pos, 2, modeTwo, relativeBase)
			three := data[pos+3]
			if one == two {
				if modeThree == 2 {
					data[three+relativeBase] = 1
				} else {
					data[three] = 1
				}
			} else {
				if modeThree == 2 {
					data[three+relativeBase] = 0
				} else {
					data[three] = 0
				}
			}
			pos = pos + 4
		}

		if opCode == 9 {
			one := getValue(data, pos, 1, modeOne, relativeBase)
			relativeBase = relativeBase + one
			pos = pos + 2
		}
	}
}

func getValue(data []int, pos, index, mode, relativeBase int) int {
	switch mode {
	case 0:
		return data[data[pos+index]]
	case 1:
		return data[pos+index]
	case 2:
		return data[data[pos+index]+relativeBase]
	default:
		return 0
	}
}
