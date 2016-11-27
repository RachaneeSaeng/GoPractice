package math //must in same package

import (
	"fmt"
	"testing"
	"testing/quick"
)

//where Xxx can be any alphanumeric string (but the first letter must not be in [a-z])
//and serves to identify the test routine.
func TestAdder(t *testing.T) {
	result := adder(4, 7)
	if result != 11 {
		t.Fatal("4 + 7 did not equal 11")
	}
}

func BenchmarkAdder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		adder(4, 7)
	}
}

//Test method startwith Example must follow coorect function name to be tested
func Exampleadder() {
	fmt.Println(adder(4, 7))
	// Output:
	// 11
}

func Exampleadder_multiple() {
	fmt.Println(adder(3, 6, 7, 4, 61))
	// Output:
	// 81
}

func TestAdderBlackbox(t *testing.T) {
	err := quick.Check(a, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func a(x, y int) bool {
	return adder(x, y) == x+y
}
