package round

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	v := []float64{
		123.456789,
		-4.5679123,
		1.23e4,
		-4.567e3,
		0x234567,
		067777777,
		123456789,
		1.234567,
		2.3456789e8,
	}

	p := []float64{
		2,
		8,
		0x3,
		04,
		3.6,
		2.1e1,
		1,
		0,
	}

	for i := 0; i < len(v); i++ {
		fmt.Printf("round: %#v\n", v[i])

		for j := 0; j < len(p); j++ {
			r := round(v[i], int(p[j]), 0)
			fmt.Printf("with precision %#v -> %#v \n", p[j], r)
		}
	}
}
