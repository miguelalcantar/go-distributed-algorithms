package les

// Leader Election in a Synchronous Ring
// Electing a unique leader process from among the processes in a network.
// A single token circulates around the network, giving its current owner the sole right
// to initiate communication.
// Sometimes, however, the token may be lost, and it becomes necessary for the processes
// to execute an algorthm to regenerate the lost token

// We assume that the network digraph G is a ring consisting of n nodes
// Node also known as process

// states of the process
type nodeLCR struct {
	uid    int
	send   *int // this way we can set send as nil
	status status

	// since process just know about their next neighbour
	next *nodeLCR
}

// By definition:
// - uid, a UID (this case int), inially i's UID
// - send, a UID or null, initially i's UID
// - status, with values in {unknown, leader}, initially unknown
func NewNodeLCR(uid int) *nodeLCR {
	return &nodeLCR{
		uid:    uid,
		send:   &uid,
		status: unknown,
	}
}

// SetNext assigns the neighbour's current node
func (n *nodeLCR) SetNext(nxt *nodeLCR) {
	n.next = nxt
}

// Msgs is the message-generation function.
// By definition: send the current value of send to process i + 1.
// If the value of the send component is null, this msgs function does not send any message.
func (n *nodeLCR) Msgs() {
	if n.send != nil {
		ChannelMap.Store(n.next.uid, n.send)
	}
}

// Trans defines the state-transition function
func (n *nodeLCR) Trans() {
	n.send = nil
	if v, ok := ChannelMap.Load(n.uid); ok {
		vint := v.(*int)
		switch {
		case *vint > n.uid:
			n.send = vint
		case *vint == n.uid:
			n.status = leader

			// send signal in order to finish the rounds cycle
			endSignal.Finish()
		}
	}
}
