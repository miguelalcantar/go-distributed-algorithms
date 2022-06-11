package les

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	endSignal = NewEndSingnal()
	os.Exit(m.Run())
}
