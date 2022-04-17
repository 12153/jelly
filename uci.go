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
var saveBM = ""

var bInfinite *bool

func uci(frGui chan string, myTell func(text ...string)) {
	*bInfinite = true
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
			handlebm(bm, bInfinite)
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
		case "ucinewgame":
			handleNewGame()
		case "position":
			handlePosition(strings.Join(words, " "))
		case "debug":
			handleDebug(words)
		case "register":
			handleRegister(words)
		case "go":
			handleGo(words)
		case "ponderhit":
			handlePonderHit()
		case "stop":
			handleStop(toEng, bInfinite)
		case "quit", "q":
			quit = true
			continue

		}
	}
}

func handlePonderHit() {

}

func handleGo(words []string) {
	if len(words) > 1 {
		words[1] = trim(low(words[1]))

		switch words[1] {
		case "searchmoves":
			tell("info string go searchmoves not implemented yet")
		case "ponder":
			tell("info string go ponder not implemented yet")
		case "wtime":
			tell("info string go wtime not implemented yet")
		case "btime":
			tell("info string go btime not implemented yet")
		case "winc":
			tell("info string go winc not implemented yet")
		case "binc":
			tell("info string go binc not implemented yet")
		case "depth":
			tell("info string go depth not implemented yet")
		case "movestogo":
			tell("info string go movestogo not implemented yet")
		case "nodes":
			tell("info string go nodes not implemented yet")
		case "movetime":
			tell("info string go movetime not implemented yet")
		case "mate":
			tell("info string go mate not implemented yet")
		case "infinite":
			tell("info string go infinite not implemented yet")
		default:
			tell("info string postion ", words[1], " not implemented")
		}
	}
}

func handleRegister(words []string) {
	tell("info string go infinite not implemented yet")

}

func handleDebug(words []string) {
	tell("info string go infinite not implemented yet")

}

func handlePosition(cmd string) {

	cmd = trim(strings.TrimPrefix(cmd, "position"))
	parts := strings.Split(cmd, "moves")
	if len(cmd) == 0 || len(parts) > 2 {
		err := fmt.Errorf("%v wrong length=%v", parts, len(parts))
		tell("info string Error", fmt.Sprint(err))
		return
	}

	alt := strings.Split(parts[0], " ")
	alt[0] = trim(alt[0])
	tell("info string position ", alt[0], "not implemented yet")

	if alt[0] == "startpos" {
		alt[0] = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	} else if alt[0] == "fen" {
		alt[0] = ""
	} else {
		err := fmt.Errorf("%#v must be fen or start pos", alt[0])
		tell("info string Error", err.Error())
		return
	}

	parts[0] = strings.Join(alt, " ")

	parts[0] = trim(parts[0])
	fmt.Printf("info string parse %#v\n", parts[0])
	parseFEN(parts[0])
	if len(parts) == 2 {
		parts[1] = low(trim(parts[1]))
		fmt.Printf("info string parse %#v\n", parts[1])
		parseMvs(parts[1])
	}
}

func handleNewGame() {
	tell("info string go infinite not implemented yet")

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

func handlebm(bm string, bInfinite *bool) {
	if *bInfinite {
		saveBM = bm
		return
	}
	tell(bm)
}

func handleStop(toEng chan string, bInfinite *bool) {

	if *bInfinite {
		if saveBM != "" {
			tell(saveBM)
			saveBM = ""
		}
	}
	toEng <- "stop"
	*bInfinite = false
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
