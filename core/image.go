package core

import (
	"image"
	_ "image/png"
	"io"
	"fmt"
)

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}


// Get the bi-dimensional pixel array
func GetPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}
func GetRGBA(file io.Reader) {
	
	img, _, err := image.Decode(file)

	if err != nil {
		fmt.Printf("%v", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			fmt.Printf("%v - %v - %v - %v\n", r, g, b, a)
			//row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
	}
	
}
// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
