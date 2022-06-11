package les

import (
	"testing"
)

func TestFM(t *testing.T) {
	n1 := NewNodeFM(5, 7)
	n2 := NewNodeFM(10, 7)
	n3 := NewNodeFM(15, 7)
	n4 := NewNodeFM(7, 7)
	n5 := NewNodeFM(20, 7)
	n6 := NewNodeFM(18, 7)
	n7 := NewNodeFM(24, 7)
	n8 := NewNodeFM(50, 7)
	n9 := NewNodeFM(100, 7)

	n1.SetOutnbrs([]*nodeFM{n2, n3})
	n2.SetOutnbrs([]*nodeFM{n3, n4, n5})
	n3.SetOutnbrs([]*nodeFM{n4, n9})
	n4.SetOutnbrs([]*nodeFM{n6})
	n5.SetOutnbrs([]*nodeFM{n6})
	n6.SetOutnbrs([]*nodeFM{n7})
	n7.SetOutnbrs([]*nodeFM{n8, n9})
	n8.SetOutnbrs([]*nodeFM{n9})

	RoundSync([]node{n1, n2, n3, n4, n5, n6})
}
