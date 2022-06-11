package les

type voids struct{}

var void voids
var globalMaxUid = 0

type nodeFM struct {
	uid    int
	maxuid int
	status status
	rounds int
	diam   int

	outnbrs []*nodeFM
}

func NewNodeFM(u, diam int) *nodeFM {
	return &nodeFM{
		uid:    u,
		diam:   diam,
		maxuid: u,
		status: unknown,
		rounds: 0,
	}
}

func (n *nodeFM) SetOutnbrs(outnbrs []*nodeFM) {
	n.outnbrs = outnbrs
}

func (n *nodeFM) Msgs() {
	for _, v := range n.outnbrs {
		a, loaded := ChannelMap.LoadOrStore(v.uid, map[int]struct{}{
			n.maxuid: void,
		})
		if loaded {
			maxuidSet := a.(map[int]struct{})
			maxuidSet[n.maxuid] = void
			ChannelMap.Store(v.uid, maxuidSet)
		}
	}
}

func (n *nodeFM) Trans() {
	n.rounds += 1

	a, loaded := ChannelMap.LoadAndDelete(n.uid)

	if loaded {
		maxuidSet := a.(map[int]struct{})
		maxuidSet[n.maxuid] = void
		n.maxuid = getMaxMapIntStruct(maxuidSet)
	}

	// extra step in order for the network to be able
	// to access to the global max state through nodes
	if n.maxuid > globalMaxUid {
		globalMaxUid = n.maxuid
	}

	if n.rounds == n.diam {
		if globalMaxUid == n.uid {
			n.status = leader
		} else {
			n.status = nonleader
		}
		// signal in order to finish the rounds cycle
		endSignal.Finish()
	}
}
