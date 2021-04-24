package main

import (
	"gocv.io/x/gocv"
	"image"
)
import "fmt"

func main() {
	filename := "baboon.jpg"
	window := gocv.NewWindow("Convolution")
	img := gocv.IMRead(filename, gocv.IMReadGrayScale)
	result := gocv.NewMat()
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}
	//kernel := gocv.Ones(3,3, gocv.MatTypeCV32FC1)
	sizes := []int{3, 3}
	s := gocv.NewScalar(3.0, -1.0, -1.0, -1.0)
	kernel := gocv.NewMatWithSizesWithScalar(sizes, gocv.MatTypeCV32FC1, s)

	gocv.Filter2D(img, &result, 3, kernel, image.Pt(-1, -1), 0 ,gocv.BorderDefault)
	for {
		window.IMShow(result)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
