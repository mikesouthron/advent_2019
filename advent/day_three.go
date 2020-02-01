package advent

import (
	"fmt"
	"strconv"
	"strings"
)

type vector struct {
	x, y, steps int
}

//ExecuteDayThree run day three
func ExecuteDayThree() {
	lines := ReadFile(3)

	wires := make(chan []vector, 2)

	go calculateWire(strings.Split(lines[0], ","), wires)
	go calculateWire(strings.Split(lines[1], ","), wires)

	firstWire := <-wires
	secondWire := <-wires

	shortest := -1

	for _, first := range firstWire {
		for _, second := range secondWire {
			if first.x == second.x && first.y == second.y {
				steps := first.steps + second.steps
				if shortest == -1 || steps < shortest {
					shortest = steps
				}
			}
		}
	}

	fmt.Println(shortest)
}

func calculateWire(dir []string, wires chan []vector) {
	x, y, steps := 0, 0, 0
	wire := make([]vector, 0)
	for _, d := range dir {
		direction := d[0]
		num, _ := strconv.Atoi(strings.TrimSpace(d[1:]))
		for i := 0; i < num; i++ {
			steps = steps + 1
			if direction == 'R' {
				x = x + 1
			}
			if direction == 'L' {
				x = x - 1
			}
			if direction == 'U' {
				y = y + 1
			}
			if direction == 'D' {
				y = y - 1
			}
			wire = append(wire, vector{x, y, steps})
		}
	}
	wires <- wire
}
