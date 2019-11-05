package aria

import "testing"

func TestAnyOverlap(t *testing.T) {
	// TODO: write test
}

func TestInexactOverlap(t *testing.T) {
	x := make([]byte, 16)
	y := x
	if inexactOverlap(x, y) {
		t.Error()
	}
}
