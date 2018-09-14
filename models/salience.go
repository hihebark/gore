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
func gaussianBlur(imgsrc image.Image) image.Image {
	bounds := imgsrc.Bounds()
	border := int(bounds.Max.Y / 2)
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	return image.Image
}
