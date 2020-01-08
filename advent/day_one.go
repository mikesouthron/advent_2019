package advent

import (
	"strconv"
	"strings"
)

func DayOneExecute() {
	lines := ReadFile(1)
	var total int64 = 0
	for _, line := range lines {
		mass, _ := strconv.Atoi(strings.TrimSpace(line))
		fuel := fuelMass(int64(mass), 0)
		total = total + fuel
	}
	println(total)
}

func fuelMass(mass, acc int64) int64 {
	newMass := (mass / 3) - 2
	if newMass <= 0 {
		return acc
	}
	return fuelMass(newMass, acc + newMass)
}



