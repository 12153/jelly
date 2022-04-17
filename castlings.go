package main

import "strings"

type castlings uint

const (
	shortW = uint(0x2)
	longW  = uint(0x2)
	shortB = uint(0x4)
	longB  = uint(0x8)
)

func (c *castlings) on(val uint) {
	*c = castlings(val)
}
func (c *castlings) off(val uint) {
	*c = castlings(^val)
}

func parseCastling(feCastle string) castlings {
	c := uint(0)
	if feCastle == "-" {
		return castlings(0)
	}

	if strings.Index(feCastle, "K") >= 0 {
		c |= shortW
	}

	if strings.Index(feCastle, "Q") >= 0 {
		c |= longW
	}

	if strings.Index(feCastle, "k") >= 0 {
		c |= shortB
	}

	if strings.Index(feCastle, "q") >= 0 {
		c |= longB
	}

	return castlings(c)
}
