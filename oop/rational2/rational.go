package rational2

import "fmt"

type Rational struct {
	n int // numerator (Zähler)
	d int // denominator (Nenner)
}

var Zero = Rational{n: 0, d: 1}
var One = Rational{n: 1, d: 1}

func New(numerator int, denominator int) Rational {
	if denominator == 0 {
		panic("denominator must not be zero")
	}
	// reduce fraction
	gcd := gcd(numerator, denominator)
	return Rational{n: numerator / gcd, d: denominator / gcd}
}

func (r Rational) Add(b Rational) Rational {
	return New(r.n*b.d+r.d*b.n, r.d*b.d)
}

func (r Rational) Multiply(b Rational) Rational {
	return New(r.n*b.n, r.d*b.d)
}

func (r Rational) String() string {
	return fmt.Sprintf("(%d/%d)", r.n, r.d)
}

func (r Rational) Float32() float32 {
	return float32(r.n) / float32(r.d)
}

func (r Rational) Float64() float64 {
	return float64(r.n) / float64(r.d)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
