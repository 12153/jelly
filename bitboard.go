package main

import "math/bits"

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
