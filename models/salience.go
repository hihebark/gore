package model

import (
	"fmt"
	"image"

	"github.com/hihebark/gore/core"
)

//Salience output salience image.
func Salience(imgsrc image.Image, kernel, radius int) image.Image {
	fmt.Printf("+ Calculating Salience map: \n")
	return core.GaussianBlur(imgsrc, kernel, radius)
}
