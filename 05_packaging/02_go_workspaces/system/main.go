package main

import (
	"fmt"

	"github.com/gabrielmq/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
}
