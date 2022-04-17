package main

import (
	"strings"
)

var trim = strings.TrimSpace
var low = strings.ToLower

func main() {
	tell("info string")
	uci(input(), mainTell)

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
