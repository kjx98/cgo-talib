package talib

import (
	"fmt"
	"math"
)

func Example() {
	fmt.Println(Sin([]float64{0, math.Pi / 2, math.Pi}))
	// Output: [0 1 1.2246467991473532e-16]
}
