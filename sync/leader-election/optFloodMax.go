package les

// node optFloodMax
type nodeOFM struct {
	uid     int
	maxuid  int
	status  status
	newinfo bool
	rounds  int
	diam    int

	outnbrs []*nodeOFM
}

func NewNodeOFM(u, diam int) *nodeOFM {
	return &nodeOFM{
		uid:     u,
		diam:    diam,
		maxuid:  u,
		status:  unknown,
		newinfo: false,
		rounds:  0,
	}
}

func (n *nodeOFM) SetOutnbrs(outnbrs []*nodeOFM) {
	n.outnbrs = outnbrs
}

func (n *nodeOFM) Msgs() {
	if n.newinfo {
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
}

func (n *nodeOFM) Trans() {
	n.rounds += 1

	U, loaded := ChannelMap.LoadAndDelete(n.uid)

	n.newinfo = false
	if loaded {
		maxU := getMaxMapIntStruct(U.(map[int]struct{}))
		if n.maxuid < maxU {
			n.maxuid = maxU
			n.newinfo = true
		}
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
