package advent

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//ReadFile Read the day file for the day int passed in
func ReadFile(day int) []string {
	data, err := ioutil.ReadFile(fmt.Sprintf("data/day%d.txt", day))
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
