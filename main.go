package main

import (
	"advent_2019/advent"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	advent.ExecuteDaySeven()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	//advent.ExecuteDayFourSync()
	//advent.ExecuteDayThree()
}
