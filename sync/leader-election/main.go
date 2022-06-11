package les

type status int

const (
	unknown status = iota
	leader
	nonleader
)

type node interface {
	Msgs()
	Trans()
}