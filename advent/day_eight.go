package advent

import (
	"fmt"
	"strconv"
)

type row struct {
	pixels []int
}

type layer struct {
	rows []row
}

//RunDayEight run day 8
func RunDayEight() {
	data := ReadFile(8)[0]
	layers := buildLayers(data)
	calculateValidLayer(layers)

	finalImage := make([][]int, 6)

	for r := 0; r < 6; r++ {
		row := make([]int, 25)
		for c := 0; c < 25; c++ {
			for _, layer := range layers {
				pixel := layer.rows[r].pixels[c]
				if pixel != 2 {
					row[c] = pixel
					break
				}
			}
		}
		finalImage[r] = row
	}

	for _, row := range finalImage {
		for _, col := range row {
			if col == 1 {
				fmt.Print(col)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func buildLayers(data string) []layer {
	layers := make([]layer, 0)

	rows := make([]row, 6)
	pixels := make([]int, 25)

	pixelCount := 0
	rowCount := 0

	for _, char := range data {
		pixel, _ := strconv.Atoi(string(char))
		pixels[pixelCount] = pixel
		pixelCount = pixelCount + 1
		if pixelCount == 25 {
			rows[rowCount] = row{pixels}
			rowCount = rowCount + 1
			pixels = make([]int, 25)
			pixelCount = 0
			if rowCount == 6 {
				layers = append(layers, layer{rows})
				rows = make([]row, 6)
				rowCount = 0
			}
		}
	}

	return layers
}

func calculateValidLayer(layers []layer) {
	minZeroCount := 0
	calculatedValue := 0

	for _, layer := range layers {
		zeroCount := 0
		oneCount := 0
		twoCount := 0
		for _, row := range layer.rows {
			for _, pixel := range row.pixels {
				switch pixel {
				case 0:
					zeroCount = zeroCount + 1
					break
				case 1:
					oneCount = oneCount + 1
					break
				case 2:
					twoCount = twoCount + 1
					break
				}
			}
		}

		if minZeroCount == 0 || zeroCount < minZeroCount {
			minZeroCount = zeroCount
			calculatedValue = oneCount * twoCount
		}
	}

	fmt.Println(calculatedValue)
}
