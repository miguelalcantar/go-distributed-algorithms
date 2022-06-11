package les

import "math"

type directionHS int

const (
	out directionHS = iota
	in
)

// bidirectional message
type bMess struct {
	from int
	to   int
}

func newBMess(from, to int) *bMess {
	return &bMess{
		from: from,
		to:   to,
	}
}

type sendHS struct {
	v         int
	direction directionHS
	h         int // phase counter
}

func newSendHS(v int, d directionHS, h int) *sendHS {
	return &sendHS{
		v:         v,
		direction: d,
		h:         h,
	}
}

type nodeHS struct {
	uid       int
	sendLeft  *sendHS
	sendRight *sendHS
	status    status
	phase     int

	left  *nodeHS
	right *nodeHS
}

func NewNodeHS(uid int, left, right *nodeHS) *nodeHS {
	return &nodeHS{
		uid:       uid,
		sendLeft:  newSendHS(uid, out, 1),
		sendRight: newSendHS(uid, out, 1),
		status:    unknown,
		phase:     0,

		left:  left,
		right: right,
	}
}

func (n *nodeHS) SetRight(nxt *nodeHS) {
	n.right = nxt
}

func (n *nodeHS) SetLeft(nxt *nodeHS) {
	n.left = nxt
}

func (n *nodeHS) Msgs() {
	if n.sendRight != nil {
		ChannelMap.Store(
			newBMess(n.uid, n.right.uid),
			n.sendRight)
	}
	if n.sendLeft != nil {
		ChannelMap.Store(
			newBMess(n.uid, n.left.uid),
			n.sendLeft)
	}
}

func (n *nodeHS) Trans() {
	n.sendRight = nil
	n.sendLeft = nil

	mLeft, _ := ChannelMap.Load(newBMess(n.left.uid, n.uid))
	mRight, _ := ChannelMap.Load(newBMess(n.right.uid, n.uid))

	if mLeft != nil {
		sMess := mLeft.(*sendHS)
		if sMess.direction == out {
			switch {
			case sMess.v > n.uid && sMess.h > 1:
				n.sendRight = newSendHS(sMess.v, out, sMess.h-1)
			case sMess.v > n.uid && sMess.h == 1:
				n.sendLeft = newSendHS(sMess.v, in, 1)
			case sMess.v == n.uid:
				n.status = leader
			}
		}
	}
	if mRight != nil {
		sMess := mRight.(*sendHS)
		if sMess.direction == out {
			switch {
			case sMess.v > n.uid && sMess.h > 1:
				n.sendLeft = newSendHS(sMess.v, out, sMess.h-1)
			case sMess.v > n.uid && sMess.h > 1:
				n.sendRight = newSendHS(sMess.v, in, 1)
			case sMess.v == n.uid:
				n.status = leader
			}
		}
	}
	if mLeft != nil {
		sMess := mLeft.(*sendHS)
		if sMess.direction == in && sMess.v != n.uid && sMess.h == 1 {
			n.sendRight = sMess
		}
	}
	if mRight != nil {
		sMess := mRight.(*sendHS)
		if sMess.direction == in && sMess.v != n.uid && sMess.h == 1 {
			n.sendLeft = sMess
		}
	}
	if mLeft != nil && mRight != nil {
		lMess := mLeft.(*sendHS)
		rMess := mRight.(*sendHS)

		if lMess.v == n.uid && lMess.direction == in && lMess.h == 1 {
			if rMess.v == n.uid && rMess.direction == in && lMess.h == 1 {
				n.phase++
				n.sendRight = newSendHS(n.uid, out, int(math.Pow(2, float64(n.phase))))
				n.sendLeft = newSendHS(n.uid, out, int(math.Pow(2, float64(n.phase))))
			}
		}
	}

}
