package round

import (
	"math"
	"testing"
)

func TestMathRound(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	_ = minInt
	var b bool
	b = epsilonEqual(-1.2, math.Ceil(-1.53))
	if !b {
		t.Fatal()
	}
	b = epsilonEqual(-1, math.Ceil(-1.5))
	if !b {
		t.Fatal()
	}

	b = epsilonEqual(2, math.Ceil(1.5))
	if !b {
		t.Fatal()
	}

	b = epsilonEqual(-2, math.Floor(-1.5))
	if !b {
		t.Fatal()
	}

	b = epsilonEqual(1, math.Floor(1.5))
	if !b {
		t.Fatal()
	}

	b = epsilonEqual(float64(minInt), math.Ceil(float64(minInt)-0.5))
	if !b {
		t.Fatal()
	}
	b = epsilonEqual(float64(minInt+1), math.Ceil(float64(minInt)+0.5))
	if !b {
		t.Fatal()
	}
	b = epsilonEqual(float64(minInt)-1.0, Round(float64(minInt)-0.6, 0, 0))
	if !b {
		t.Fatal(b)
	}

	b = epsilonEqual(float64(minInt), Round(float64(minInt)-0.4, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(minInt), Round(float64(minInt)+0.4, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(minInt)+1, Round(float64(minInt)+0.6, 0, 0))
	if !b {
		t.Fatal(b)
	}

	b = epsilonEqual(float64(minInt)-1, math.Floor(float64(minInt)-0.5))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(minInt), math.Floor(float64(minInt)+0.5))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt), math.Ceil(float64(maxInt)-0.5))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt)+1, math.Ceil(float64(maxInt)+0.5))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt)-1, Round(float64(maxInt)-0.6, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt), Round(float64(maxInt)-0.4, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt), Round(float64(maxInt)+0.4, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt)+1, Round(float64(maxInt)+0.6, 0, 0))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt)-1, math.Floor(float64(maxInt)-0.5))
	if !b {
		t.Fatal(b)
	}
	b = epsilonEqual(float64(maxInt), math.Floor(float64(maxInt)+0.5))
	if !b {
		t.Fatal(b)
	}

}

func epsilonEqual(left, right float64) bool {
	return (math.Abs(left-right) / left) < 1e-12
}
