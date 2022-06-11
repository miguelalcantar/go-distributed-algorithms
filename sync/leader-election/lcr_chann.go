package les

type nodeLCRc struct {
	uid    int
	send   *int // this way we can set send as nil
	status status
	c      chan *int

	// since process just know about their next neighbour
	next *nodeLCRc
}

func NewNodeLCRc(uid int) *nodeLCRc {
	return &nodeLCRc{
		uid:    uid,
		send:   &uid,
		status: unknown,
		c:      make(chan *int),
	}
}

func (n *nodeLCRc) SetNext(nxt *nodeLCRc) {
	n.next = nxt
}

func (n *nodeLCRc) Msgs() {
	go func() {
		if n.send != nil {
			n.next.c <- n.send
		}
	}()
}

func (n *nodeLCRc) Trans() {
	send := <-n.c
	switch {
	case *send > n.uid:
		n.send = send
	case *send == n.uid:
		n.status = leader

		// send signal in order to finish the rounds cycle
		endSignal.Finish()
		close(n.c) // closing channel
	}
}
