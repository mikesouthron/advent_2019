package advent

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(day int) []string {
	data, err := ioutil.ReadFile(fmt.Sprintf("data/day%d.txt", day))
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}