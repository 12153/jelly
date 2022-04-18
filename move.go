package main

// directions
const (
	E           = +1
	W           = -1
	N           = 8
	S           = -8
	NW          = +7
	NE          = +9
	SW          = -NE
	SE          = -NW
	toShift     = 6
	p12Shift    = 6 + toShift
	cpShift     = 4 + p12Shift
	prShift     = 4 + cpShift
	epShift     = 4 + prShift
	castleShift = 6 + epShift
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

func (m *move) packMove(fr, to, p12, cp, pr, ep uint, castl castlings) {
	// 6 bits frm, 6 bits to, 4 bits piece(p12), 4 bits captured piece, 4 bits pr, 6 bits ep, 4 bits castle,  x bits value
	*m = move(fr | (to << toShift) | (p12 | p12Shift) | (cp << cpShift) | (pr << prShift) | (ep << epShift) | (uint(castl) << castleShift))
}
