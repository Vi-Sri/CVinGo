package main

import "gonum.org/v1/gonum/mat"
import "fmt"
import "math/rand"

func main() {
	d := mat.NewDense(3, 4, nil)

	// Set d to have random integer elements.
	m, n := d.Dims()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			d.Set(i, j, float64(rand.Intn(10)))
		}
	}
	fmt.Printf("%v\n", mat.Formatted(d))

	// Print the transpose.
	fmt.Printf("%v\n", mat.Formatted(d.T()))

	a := mat.NewDense(2, 3, []float64{
		3, 4, 5,
		1, 2, 3,
	})
	b := mat.NewDense(3, 3, []float64{
		1, 1, 8,
		1, 2, -3,
		5, 5, 7,
	})
	fmt.Println("trace(b) =", mat.Trace(b))

	c := &mat.Dense{}
	c.Mul(a, b)
	c.Add(c, a)
	c.Mul(c, b.T())
	fmt.Printf("%v\n", mat.Formatted(c))
}
