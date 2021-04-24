package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {

	filename := "baboon.jpg"
	window := gocv.NewWindow("Hello world!!")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}
	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}