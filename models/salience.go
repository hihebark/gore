package model

import (
	"fmt"
	"image"
)

//Salience output salience image.
func Salience(imgsrc image.Image) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	return imgsrc
}
