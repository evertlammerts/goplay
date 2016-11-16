package integral

import (
	"math"
	"testing"
)

func TestTrapezoid(t *testing.T) {
	for i, test := range []struct {
		numStrips                  int64
		intervalStart, intervalEnd float64
		f                          func(x float64) float64
		want                       float64
	}{
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   1,
			f:             func(x float64) float64 { return x },
			want:          0.5,
		},
		{
			numStrips:     1e6,
			intervalStart: -1,
			intervalEnd:   1,
			f:             func(x float64) float64 { return x },
			want:          0,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   1,
			f:             func(x float64) float64 { return x + 10 },
			want:          10.5,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   1,
			f:             func(x float64) float64 { return 3*x*x + 10 },
			want:          11,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   1,
			f:             func(x float64) float64 { return math.Exp(x) },
			want:          1.7182818284591876,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   math.Pi,
			f:             func(x float64) float64 { return math.Cos(x) },
			want:          0,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   2 * math.Pi,
			f:             func(x float64) float64 { return math.Cos(x) },
			want:          0,
		},
		{
			numStrips:     1e7,
			intervalStart: 0,
			intervalEnd:   math.Pi,
			f:             func(x float64) float64 { return math.Sin(x) },
			want:          2,
		},
		{
			numStrips:     1e7,
			intervalStart: 0,
			intervalEnd:   0.5 * math.Pi,
			f:             func(x float64) float64 { return math.Sin(x) },
			want:          1,
		},
		{
			numStrips:     1e6,
			intervalStart: 0,
			intervalEnd:   2 * math.Pi,
			f:             func(x float64) float64 { return math.Sin(x) },
			want:          0,
		},
	} {
		v := Trapezoid(test.numStrips, test.intervalStart, test.intervalEnd, test.f)
		if v != test.want && math.Abs(test.want-v) > 1e-10 {
			t.Errorf("test #%d: got=%v want=%v\n", i, v, test.want)
		}
	}
}
