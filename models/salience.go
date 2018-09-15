package model

import (
	"fmt"
	"image"
)

//Salience output salience image.
func Salience(imgsrc image.Image, gaussKernel int) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	return imgsrc
}

func gaussianBlur(imgsrc image.Image, gaussKernel int) image.Image {
	bounds := imgsrc.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			_ := Gaussian(x, y, float64(5)) // to end this
		}
	}
	return image.Image
}
