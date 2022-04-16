package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var tell = mainTell
var trim = strings.TrimSpace
var low = strings.ToLower

func uci(frGui chan string, myTell func(text ...string)) {
	tell = myTell
	tell("hello from uci")
	quit := false
	cmd := ""
	bm := ""
	words := []string{}
	frEng, toEng := engine()

	for quit == false {
		select {
		case cmd := <-frGui:
			words = strings.Split(cmd, " ")
		case bm = <-frEng:
			handlebm(bm)
			continue
		}

		words[0] = trim(words[0])
		switch words[0] {
		case "uci":
			handleUci()
		case "isready":
			handleIsReady()
		case "setoption":
			handleSetOption([]string{cmd})
		case "stop":
			handleStop(toEng)
		case "quit", "q":
			quit = true
			continue

		}
	}
}

func handleUci() {
	tell("id name jelly")
	tell("id author fish")

	tell("option name Hash type spin default 32 min 1 max 1024")
	tell("option name Threads type spin default 1 min 1 max 16")
	tell("uciok")

}
func handleIsReady() {
	tell("readyok")
}
func handleSetOption(option []string) {
	tell("info string set option", strings.Join(option, " "))
	tell("info string not implemented yet")
}

func handlebm(bm string) {
	tell(bm)
}

func handleStop(toEng chan string) {
	toEng <- "stop"
}

func input() chan string {
	line := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()

	return line
}

func mainTell(text ...string) {

	toGui := ""
	for _, v := range text {
		toGui += v
	}

	fmt.Println(toGui)
}
