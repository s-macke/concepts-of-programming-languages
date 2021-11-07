package forth

import (
	"strings"
	"testing"
)

func assertEqual(t *testing.T, left, right string, err error) {
	right = strings.TrimSpace(right)
	left = strings.TrimSpace(left)
	if err != nil {
		t.Errorf("assert failed with error: %v", err)
	}
	if left != right {
		t.Errorf("assert failed with: '%s' is not equal to '%s'", left, right)
	}
}

func TestMath(t *testing.T) {
	f := NewForth()
	result, err := f.Exec("1 2 + .")
	assertEqual(t, result, "3", err)

	result, err = f.Exec("1 2 - .")
	assertEqual(t, result, "-1", err)

	result, err = f.Exec("3 2 * .")
	assertEqual(t, result, "6", err)

	result, err = f.Exec("6 2 / .")
	assertEqual(t, result, "3", err)
}

func TestVariable(t *testing.T) {
	f := NewForth()

	result, err := f.Exec("variable var1")
	assertEqual(t, result, "", err)

	result, err = f.Exec("100 var1 !")
	assertEqual(t, result, "", err)

	result, err = f.Exec("var1 @ .")
	assertEqual(t, result, "100", err)
}

func TestConstants(t *testing.T) {
	f := NewForth()
	result, err := f.Exec("100 constant const1")
	assertEqual(t, result, "", err)

	result, err = f.Exec("const1 .")
	assertEqual(t, result, "100", err)
}

func TestCompare(t *testing.T) {
	f := NewForth()
	result, err := f.Exec("5 5 = .")
	assertEqual(t, result, "-1", err)

	result, err = f.Exec("4 3 = .")
	assertEqual(t, result, "0", err)

	result, err = f.Exec("6 5 > .")
	assertEqual(t, result, "-1", err)

	result, err = f.Exec("6 5 < .")
	assertEqual(t, result, "0", err)

	result, err = f.Exec("5 5 <> .")
	assertEqual(t, result, "-1", err)
}

func TestComment(t *testing.T) {
	f := NewForth()
	result, err := f.Exec("1 ( 2 3 then loop do this is my random comment ) 4 5 .s")
	assertEqual(t, result, "<3> 1 4 5", err)
}

func TestIfThen(t *testing.T) {
	f := NewForth()
	result, err := f.Exec("1 if 2 . then")
	assertEqual(t, result, "2", err)

	result, err = f.Exec("-1 if 3 . then")
	assertEqual(t, result, "3", err)

	result, err = f.Exec("0 if 2 . then")
	assertEqual(t, result, "", err)
}

func TestFunc(t *testing.T) {
	f := NewForth()
	result, err := f.Exec(`: square dup * ;
                              3 square .`)
	assertEqual(t, result, "9", err)
}

func TestFuncIfThen(t *testing.T) {
	f := NewForth()
	result, err := f.Exec(": test 1 if 2 . then 3 . ;")
	result, err = f.Exec("test")
	assertEqual(t, result, "2 3", err)
}

func TestRecurse(t *testing.T) {
	f := NewForth()
	result, err := f.Exec(": rec 1 - dup . dup 0 > if recurse then ;")
	result, err = f.Exec("4 rec")
	assertEqual(t, result, "3 2 1 0", err)
}

func TestLoop(t *testing.T) {
	f := NewForth()
	result, err := f.Exec(": iter 5 0 do i . loop ;")
	result, err = f.Exec("iter")
	assertEqual(t, result, "0 1 2 3 4", err)
}

/*
func TestFaculty(t *testing.T) {
    f := NewForth()
    result := f.Exec(`
        : fac2
        dup 0> if
        dup 1- recurse *
        else
        drop 1
        endif ;
        8 fac2 .
    `)
}
*/
