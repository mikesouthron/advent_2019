package advent

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

//ExecuteDayFour run day 4
//Go routine version is much much slower, go routine overhead is a big drain when the function is fast!
func ExecuteDayFour() {
	var wg sync.WaitGroup
	var count int64
	for i := 171309; i < 643603; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s := strconv.Itoa(i)
			if match(s) {
				atomic.AddInt64(&count, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(count)
}

//ExecuteDayFourSync run day 4 with no go routines
func ExecuteDayFourSync() {
	matches := 0
	for i := 171309; i < 643603; i++ {
		s := strconv.Itoa(i)
		if match(s) {
			matches = matches + 1
		}
	}
	fmt.Println(matches)
}

func match(s string) bool {
	match := true
	anyDoubles := false
	length := len(s)
	previous := make([]uint8, 0)
	for j := 0; j < length; j++ {
		current := s[j]
		if len(previous) > 0 && previous[len(previous)-1] != current {
			if len(previous) == 2 {
				anyDoubles = true
			}
			previous = make([]uint8, 0)
		}
		if j != length-1 {
			if s[j] > s[j+1] {
				match = false
				break
			}
		}
		previous = append(previous, current)
	}
	if len(previous) == 2 {
		anyDoubles = true
	}
	return match && anyDoubles
}
