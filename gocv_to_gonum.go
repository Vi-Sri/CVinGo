package main

import (
	"fmt"

	"gocv.io/x/gocv"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// read image as grayscale
	img := gocv.IMRead("baboon.jpg", gocv.IMReadGrayScale)
	fmt.Printf("this image have %d rows and %d cols and have the type: %s\n", img.Rows(), img.Cols(), img.Type())
	// this image have 1080 rows and 1920 cols and have the type: %!s(gocv.MatType=0)

	// convert image to float64
	convertedImg := gocv.NewMat()
	img.ConvertTo(&convertedImg, gocv.MatTypeCV64F)
	imgArray, err := convertedImg.DataPtrFloat64()
	if err != nil {
		fmt.Printf("Error converting image to float64: %v\n", err)
		return
	}

	fmt.Printf("imgArray length: %d == %d\n", len(imgArray), img.Rows()*img.Cols())
	// imgArray length: 2073600 == 2073600

	// create a new mat.Dense
	dense := mat.NewDense(img.Rows(), img.Cols(), imgArray)

	rows, cols := dense.Dims()
	fmt.Printf("Dense: %d %d\n", rows, cols)
	// Dense: 1080 1920
}