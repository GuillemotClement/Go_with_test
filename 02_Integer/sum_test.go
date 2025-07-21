package integer

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(2, 2)
	want := 4

	if sum != want {
		t.Errorf("got %d, want %d", sum, want)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
