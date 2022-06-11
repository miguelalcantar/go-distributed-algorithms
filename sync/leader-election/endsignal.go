package les

import "sync"

var (
	once      sync.Once
	endSignal *EndSignal
)

type EndSignal struct {
	mu  sync.Mutex
	end bool
}

func NewEndSingnal() *EndSignal {
	return &EndSignal{
		end: false,
	}
}

// Singleton thread safe
func GetSingleInstanceEndSignal() *EndSignal {
	if endSignal == nil {
		once.Do(
			func() {
				endSignal = NewEndSingnal()
			})
	}
	return endSignal
}

// Finish sets the end variable to true thread safetly
func (e *EndSignal) Finish() {
	e.mu.Lock()
	e.end = true
	e.mu.Unlock()
}
