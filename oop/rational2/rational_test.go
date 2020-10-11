package rational2

import "testing"

func TestString(t *testing.T) {
	var s string

	s = New(2, 3).String()
	if s != "(2/3)" {
		t.Errorf("expected 2/3, got %s", s)
	}

	s = New(2, -3).String()
	if s != "(-2/3)" {
		t.Errorf("expected -2/3, got %s", s)
	}

	s = New(12, 4).String()
	if s != "(3/1)" {
		t.Errorf("expected 3/1, got %s", s)
	}

	s = New(0, 4).String()
	if s != "(0/1)" {
		t.Errorf("expected 0/1, got %s", s)
	}
}

func TestGcd(t *testing.T) {
	var g int

	g = gcd(2, 3)
	if g != 1 {
		t.Errorf("gcd(2,3) should be 1, but was %d", g)
	}

	g = gcd(9, 3)
	if g != 3 {
		t.Errorf("gcd(9,3) should be 3, but was %d", g)
	}

	g = gcd(-2, 4)
	if g != -2 {
		t.Errorf("gcd(-2,4) should be -2, but was %d", g)
	}
}

func TestAdd(t *testing.T) {
	a := New(2, 3)
	b := New(1, 2)
	c := New(1, 3)

	s1 := a.Add(b)
	if s1 != New(7, 6) {
		t.Errorf("2/3 + 1/2 should be 7/6, but was %v", s1)
	}

	s2 := b.Add(c)
	if s2 != New(5, 6) {
		t.Errorf("1/2 + 1/3 should be 5/6, but was %v", s2)
	}

	s3 := a.Add(c)
	if s3 != New(1, 1) {
		t.Errorf("2/3 + 1/3 should be 1/1, but was %v", s3)
	}

	s4 := a.Add(Zero)
	if s4 != a {
		t.Errorf("2/3 + 0 should be 2/3, but was %v", s4)
	}
}

func TestMultiply(t *testing.T) {
	a := New(2, 3)
	b := New(1, 2)
	c := New(1, 3)

	s1 := a.Multiply(b)
	if s1 != New(2, 6) {
		t.Errorf("2/3 * 1/2 should be 2/6, but was %v", s1)
	}

	s2 := b.Multiply(c)
	if s2 != New(1, 6) {
		t.Errorf("1/2 * 1/3 should be 1/6, but was %v", s2)
	}

	s3 := a.Multiply(c)
	if s3 != New(2, 9) {
		t.Errorf("2/3 * 1/3 should be 2/9, but was %v", s3)
	}

	s4 := a.Multiply(One)
	if s4 != a {
		t.Errorf("2/3 * 1 should be 2/3, but was %v", s4)
	}

	s5 := a.Multiply(Zero)
	if s5 != Zero {
		t.Errorf("2/3 * 0 should be 0/1, but was %v", s5)
	}
}
