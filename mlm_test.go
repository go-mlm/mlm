package mlm

import (
	"testing"
)

func Test(t *testing.T) {
	mlm := NewMlm("test")
	t.Logf("%#+v", mlm)

	mlm.R()
}
