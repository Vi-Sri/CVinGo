package main

import "gonum.org/v1/gonum/mat"
import "fmt"

func main() {
	d := mat.NewDense(3, 4, nil)
	fmt.Printf("%v\n", mat.Formatted(d))

	// Initialize a matrix with pre-allocated data. Data has row-major storage.
	data := []float64{
		6, 3, 5,
		-1, 9, 7,
		2, 3, 4,
	}
	d2 := mat.NewDense(3, 3, data)
	fmt.Printf("%v\n", mat.Formatted(d2))

}
