package core

import (
	"image"
	"image/png"
	"image/color"
	"io"
	"fmt"
	"os"
)

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func MakeItGray(i io.Reader){
	src, _, err := image.Decode(i)
	if err != nil {
		fmt.Printf("%v\n", err)
    }
    bounds := src.Bounds()
    w, h := bounds.Max.X, bounds.Max.Y
    gray := image.NewGray(bounds)
    for x := 0; x < w; x++ {
        for y := 0; y < h; y++ {
            oldColor := src.At(x, y)
            grayColor := color.GrayModel.Convert(oldColor)
            gray.Set(x, y, grayColor)
        }
    }

    // Encode the grayscale image to the output file
    outfile, err := os.Create("data/gray.png")
    if err != nil {
    	fmt.Printf("%v\n", err)
    }
    defer outfile.Close()
    png.Encode(outfile, gray)
}

// Get the bi-dimensional pixel array
func GetPixels(file io.Reader) ([][]Pixel, error) {
	img, format, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Printf("image Format: %s\n", format)
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
// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
