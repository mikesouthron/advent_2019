package main

import (
	"advent_2019/advent"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	advent.RunDayNine()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//advent.ExecuteDayFourSync()
	//advent.ExecuteDayThree()
}
