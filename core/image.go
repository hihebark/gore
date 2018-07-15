package core

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
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
		fmt.Printf("image:Start:os.open base Image image:%s", i)
	}
	n, err := img.Stat()
	fmt.Printf("Converting %s to gray.\n", n.Name())
	makeItGray(img, n.Name())
	imggray, err := os.Open(fmt.Sprintf("data/gray-%s", n.Name()))
	defer imggray.Close()
	if err != nil {
		fmt.Printf("image:Start:os.open grayImage image:%s", i)
	}
	checkPixel(imggray, n.Name())
	//	pixels, err := getPixels(imggray)
	//	if err != nil {
	//		fmt.Printf("image:Start:getPixels: image Format %v", err)
	//	}
	//	fmt.Printf("%v\n", pixels)
}

func makeItGray(i io.Reader, n string) {
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
	outfile, err := os.Create(fmt.Sprintf("data/gray-%s", n))
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer outfile.Close()
	png.Encode(outfile, gray)
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
			arrow.Set(x, y, color.RGBA{255, 255, 255, 255})
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
			switch position[m.niv][m.key] {
			case "upleft":
				arrow.Set(v-1, z-1, color.RGBA{255, 0, 0, 255})
			case "up":
				arrow.Set(v, z-1, color.RGBA{255, 0, 0, 255})
			case "upright":
				arrow.Set(v+1, z-1, color.RGBA{255, 0, 0, 255})
			case "right":
				arrow.Set(v+1, z, color.RGBA{255, 0, 0, 255})
			case "left":
				arrow.Set(v-1, z, color.RGBA{255, 0, 0, 255})
			case "downleft":
				arrow.Set(v-1, z+1, color.RGBA{255, 0, 0, 255})
			case "down":
				arrow.Set(v, z+1, color.RGBA{255, 0, 0, 255})
			case "downright":
				arrow.Set(v+1, z+1, color.RGBA{255, 0, 0, 255})
			}
			arrow.Set(x, y, color.RGBA{255, 0, 0, 255})
		}
	}
	outfile, err := os.Create(fmt.Sprintf("data/gray-arrow-%s", n))
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
