package main

import (
	"fmt"
	"math/bits"
	"strings"
)

type bitboard uint64

func (b bitboard) count() int {
	return bits.OnesCount64(uint64(b))
}

func (b *bitboard) set(pos uint) {
	*b |= bitboard(uint64(1) << pos)
}

func (b bitboard) test(pos uint) bool {
	return (b & bitboard(uint(1)<<pos)) != 0
}

func (b *bitboard) clr(pos uint) {
	*b &= bitboard(^(uint64(1) << pos))
}

func (b *bitboard) firstOne() int {
	bit := bits.TrailingZeros64(uint64(*b))
	if bit == 64 {
		return 64
	}
	*b = (*b >> uint(bit+1)) << uint(bit+1)
	return bit
}

func (b bitboard) String() string {
	zeroes := ""
	for i := 0; i < 64; i++ {
		zeroes += "0"
	}

	bits := zeroes + fmt.Sprintf("%b", b)
	return bits[len(bits)-64:]
}

func (b bitboard) Stringln() string {
	s := b.String()
	row := [8]string{}
	row[0] = s[:8]
	row[1] = s[8:16]
	row[2] = s[16:24]
	row[3] = s[24:32]
	row[4] = s[32:40]
	row[5] = s[40:48]
	row[6] = s[48:56]
	row[7] = s[56:]
	for i, r := range row {
		row[i] = Reverse(r)
	}
	return strings.Join(row[:], " ")
}

func Reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}
