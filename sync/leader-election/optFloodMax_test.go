package les

import "testing"

func TestOFM(t *testing.T) {
	n1 := NewNodeOFM(5, 7)
	n2 := NewNodeOFM(10, 7)
	n3 := NewNodeOFM(15, 7)
	n4 := NewNodeOFM(7, 7)
	n5 := NewNodeOFM(20, 7)
	n6 := NewNodeOFM(18, 7)
	n7 := NewNodeOFM(24, 7)
	n8 := NewNodeOFM(50, 7)
	n9 := NewNodeOFM(100, 6)

	n1.SetOutnbrs([]*nodeOFM{n2, n3})
	n2.SetOutnbrs([]*nodeOFM{n3, n4, n5})
	n3.SetOutnbrs([]*nodeOFM{n4, n9})
	n4.SetOutnbrs([]*nodeOFM{n6})
	n5.SetOutnbrs([]*nodeOFM{n6})
	n6.SetOutnbrs([]*nodeOFM{n7})
	n7.SetOutnbrs([]*nodeOFM{n8, n9})
	n8.SetOutnbrs([]*nodeOFM{n9})

	RoundSync([]node{n1, n2, n3, n4, n5, n6})

	// leader - nonleader assertion
}
