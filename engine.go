package main

import "fmt"

func engine() (frEng, toEng chan string) {
	fmt.Println("HEllo from engine")

	frEng = make(chan string)
	toEng = make(chan string)

	go func() {
		for cmd := range toEng {
			switch cmd {
			case "stop":
			case "quit":
			}
		}
	}()
	return
}
