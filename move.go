package main

// directions
const (
	E  = +1
	W  = -1
	N  = 8
	S  = -8
	NW = +7
	NE = +9
	SW = -NE
	SE = -NW
)

var piecRules [nP][]int // not pawns

func init() {
	piecRules[ROOK] = append(piecRules[ROOK], N)
	piecRules[ROOK] = append(piecRules[ROOK], E)
	piecRules[ROOK] = append(piecRules[ROOK], S)
	piecRules[ROOK] = append(piecRules[ROOK], W)

	piecRules[BISHOP] = append(piecRules[BISHOP], NW)
	piecRules[BISHOP] = append(piecRules[BISHOP], NE)
	piecRules[BISHOP] = append(piecRules[BISHOP], SW)
	piecRules[BISHOP] = append(piecRules[BISHOP], SE)
}

type MoveList struct {
	mv []move
}

type move uint64

func (ml *MoveList) add(mv move) {
	ml.mv = append(ml.mv, mv)
}

var ml = MoveList{}
