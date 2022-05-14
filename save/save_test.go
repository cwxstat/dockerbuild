package save

import (
	"testing"
)

func TestInit(t *testing.T) {
	_, err := Init("/tmp/stuff")
	if err != nil {
		t.FailNow()
	}
}
