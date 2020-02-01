package advent

import (
	"fmt"
	"strings"
	"sync/atomic"
)

//ExecuteDaySix run day six
func ExecuteDaySix() {
	lines := ReadFile(6)
	orbits := make(map[string]string)

	for _, line := range lines {
		orbit := strings.Split(strings.TrimSpace(line), ")")
		orbits[orbit[1]] = orbit[0]
	}

	count(orbits)

	you := listOrbitsFrom(orbits, "YOU", make([]string, 0))
	san := listOrbitsFrom(orbits, "SAN", make([]string, 0))

	done := false

	for i, y := range you {
		for j, s := range san {
			if y == s {
				done = true
				fmt.Println(i + j)
				break
			}
		}
		if done {
			break
		}
	}
}

func count(orbits map[string]string) {
	var count int32 = 0
	for _, v := range orbits {
		atomic.AddInt32(&count, 1)
		countOrbitsFrom(orbits, v, &count)
	}
	fmt.Println(count)
}

func countOrbitsFrom(orbits map[string]string, start string, count *int32) {
	if val, ok := orbits[start]; ok {
		atomic.AddInt32(count, 1)
		countOrbitsFrom(orbits, val, count)
	}
}

func listOrbitsFrom(orbits map[string]string, start string, list []string) []string {
	if val, ok := orbits[start]; ok {
		list = append(list, val)
		return listOrbitsFrom(orbits, val, list)
	}
	return list
}
