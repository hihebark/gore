package core

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
)

type pixel struct {
	R int
	G int
	B int
	A int
}
type rect struct {
	up        color.Gray
	down      color.Gray
	right     color.Gray
	left      color.Gray
	center    color.Gray
	upleft    color.Gray
	upright   color.Gray
	downleft  color.Gray
	downright color.Gray
}
type maxArray struct {
	val uint8
	key int
	niv int
}

func Start(i string) {
	img, err := os.Open(i)
	defer img.Close()
	if err != nil {
		fmt.Printf("image:Start:os.open base Image image:%s\n", i)
	}
	n, _ := img.Stat()
	name := strings.Split(n.Name(), ".")[0]
	imdecode := decode(img)
	if err != nil {
		fmt.Printf("error image decode image: %s error: %v\n", name, err.Error)
	}
	if imdecode.ColorModel() != color.GrayModel {
		fmt.Printf("Converting image to grayscale\n")
		makeItGray(imdecode, name)
	}
	imggray, err := os.Open(fmt.Sprintf("data/gray-%s.gore.png", name))
	defer imggray.Close()
	if err != nil {
		fmt.Printf("image:Start:os.open grayImage image:%s\n", i)
	}
	im := decode(imggray)
	p := image.Point{X: 0, Y: 0}
	DrawSB(p, im)
	//checkPixel(imggray, name)
	//	pixels, err := getPixels(imggray)
	//	if err != nil {
	//		fmt.Printf("image:Start:getPixels: image Format %v", err)
	//		fmt.Printf("image:Start:getPixels: image Format %v", err)
	//	}
	//	fmt.Printf("%v\n", pixels)
}

func decode(i io.Reader) image.Image {
	img, _, err := image.Decode(i)
	if err != nil {
		fmt.Printf("error decode image : %v\n", err)
		//return nil, err
		panic("Decode")
	}
	return img
}

func makeItGray(img image.Image, n string) {

	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	// Encode the grayscale image to the output file
	outfile, err := os.Create(fmt.Sprintf("data/gray-%s.gore.png", n))
	if err != nil {
		fmt.Printf("image.go:makeItGray:os.Create: image: %s %v\n", n, err)
	}
	defer outfile.Close()
	png.Encode(outfile, gray)
}

func splitImagetoSquare(img image.Image) {
	//split my image to 16x16 sub image and
}

func checkPixel(i io.Reader, n string) {
	img, _, err := image.Decode(i)
	if err != nil {
		fmt.Printf("image:checkPixel: %v\n", err)
	}
	bounds := img.Bounds()
	arrow := image.NewGray(bounds)
	width, height := bounds.Max.X, bounds.Max.Y
	position := [][]string{
		{"upleft", "up", "upright"},
		{"left", "center", "right"},
		{"downleft", "down", "downright"},
	}
	m := maxArray{
		key: 0,
		val: 0,
		niv: 0,
	}
	r := rect{}
	ar := [][]uint8{}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			arrow.Set(x, y, color.White)
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v, z := x, y
			r = rect{
				up:        color.GrayModel.Convert(img.At(v, z-1)).(color.Gray),
				down:      color.GrayModel.Convert(img.At(v, z+1)).(color.Gray),
				right:     color.GrayModel.Convert(img.At(v+1, z)).(color.Gray),
				left:      color.GrayModel.Convert(img.At(v-1, z)).(color.Gray),
				center:    color.GrayModel.Convert(img.At(x, y)).(color.Gray),
				upleft:    color.GrayModel.Convert(img.At(v-1, z-1)).(color.Gray),
				upright:   color.GrayModel.Convert(img.At(v+1, z-1)).(color.Gray),
				downright: color.GrayModel.Convert(img.At(v+1, z+1)).(color.Gray),
				downleft:  color.GrayModel.Convert(img.At(v-1, z+1)).(color.Gray),
			}
			ar = [][]uint8{
				{r.upleft.Y, r.up.Y, r.upright.Y},
				{r.left.Y, 0, r.right.Y},
				{r.downleft.Y, r.down.Y, r.downright.Y},
			}

			for k, v := range ar {
				for key, val := range v {
					if val > m.val {
						m = maxArray{
							key: key,
							val: val,
							niv: k,
						}
					}
				}
			}
			//fmt.Printf("%s - ",position[m.niv][m.key])
			switch position[m.key][m.niv] {
			case "upleft":
				arrow.Set(v-1, z-1, color.Black)
			case "up":
				arrow.Set(v, z-1, color.Black)
			case "upright":
				arrow.Set(v+1, z-1, color.Black)
			case "right":
				arrow.Set(v+1, z, color.Black)
			case "left":
				arrow.Set(v-1, z, color.Black)
			case "downleft":
				arrow.Set(v-1, z+1, color.Black)
			case "down":
				arrow.Set(v, z+1, color.Black)
			case "downright":
				arrow.Set(v+1, z+1, color.Black)
			}
			arrow.Set(x, y, color.Black)
		}
	}
	outfile, err := os.Create(fmt.Sprintf("data/gray-arrow-%s.gore.png", n))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer outfile.Close()
	png.Encode(outfile, arrow)
	fmt.Printf("ar: %v\nm: %v\nr: %v\n", ar, m, r)
}

// Get the bi-dimensional pixel array
func getPixels(i io.Reader) ([][]pixel, error) {
	img, format, err := image.Decode(i)
	if err != nil {
		return nil, err
	}
	fmt.Printf("image Format: %s\n", format)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]pixel
	for y := 0; y < height; y++ {
		var row []pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel {
	return pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
