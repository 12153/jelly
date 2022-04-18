package main

import (
	"fmt"
	"math"
)

type searchLimits struct {
	Depth    int
	Nodes    uint64
	MoveTime int // milliseconds
	Infinite bool
}

var limits searchLimits

func (s *searchLimits) init() {
	s.Depth = 9999
	s.Nodes = math.MaxUint64
	s.MoveTime = 99999999
	s.Infinite = false
}

func engine() (frEng, toEng chan string) {
	fmt.Println("HEllo from engine")

	frEng = make(chan string)
	toEng = make(chan string)

	go func() {
		for cmd := range toEng {
			switch cmd {
			case "stop":
			case "quit":
			case "go":
				tell("info string i'm thinking")
			}
		}
	}()
	return
}

func (s *searchLimits) SetDepth(int) {
	s.Depth = 9999

}
func (s *searchLimits) setMoveTime(val int) {
	s.MoveTime = val
}
func (s *searchLimits) setNodes(val uint64) {
	s.Nodes = val
}
func (s *searchLimits) setInfinite(val bool) {
	s.Infinite = val
}
