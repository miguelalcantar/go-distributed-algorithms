package les

import (
	"fmt"
	"testing"
)

func TestLcr(t *testing.T) {
	n1 := NewNodeLCR(123)
	n2 := NewNodeLCR(98)
	n3 := NewNodeLCR(75)
	n4 := NewNodeLCR(15)
	n5 := NewNodeLCR(512)
	n6 := NewNodeLCR(51)

	n1.SetNext(n2)
	n2.SetNext(n3)
	n3.SetNext(n4)
	n4.SetNext(n5)
	n5.SetNext(n6)
	n6.SetNext(n1)

	RoundSync([]node{n1, n2, n3, n4, n5, n6})

	//assert leader
	fmt.Print("")
}

func TestLCRc(t *testing.T) {
	n1 := NewNodeLCRc(123)
	n2 := NewNodeLCRc(98)
	n3 := NewNodeLCRc(75)
	n4 := NewNodeLCRc(15)
	n5 := NewNodeLCRc(512)
	n6 := NewNodeLCRc(51)

	n1.SetNext(n2)
	n2.SetNext(n3)
	n3.SetNext(n4)
	n4.SetNext(n5)
	n5.SetNext(n6)
	n6.SetNext(n1)

	RoundSync([]node{n1, n2, n3, n4, n5, n6})

	//assert leader
	fmt.Print("")
}