package tempconv

import (
	"fmt"
	"testing"
)

func TestKelvin(t *testing.T) {
	k := Kelvin(285.0)

	fmt.Printf("Celsius: %.3f from kelvin [%.3f]\n", KToC(k), k)
}
