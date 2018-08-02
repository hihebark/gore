package core

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

type imageInfo struct {
	format string
	name   string
	bounds image.Rectangle
}

func newImageInfo(f, n string, b image.Rectangle) *imageInfo {
	return &imageInfo{
		format: f,
		name:   n,
		bounds: b,
	}
}

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

func Start(path string) {

	img, err := os.Open(path)
	defer img.Close()
	if err != nil {
		fmt.Printf("image:start:os.Open path:%s\n", path)
	}

	info, _ := img.Stat()
	name := strings.Split(info.Name(), ".")[0]
	imgdec, form := decode(img)
	ii := newImageInfo(form, name, imgdec.Bounds())
	var gray image.Image
	if imgdec.ColorModel() != color.GrayModel {
		gray = ii.grayscaleI(imgdec)
	}
	//func DrawSB(p image.Point, img image.Image) image.Image
	//square := DrawSB(image.Pt(100, 100), gray)
	//func DrawLine(start, end image.Point, img image.Image, thick int, c color.Color)
	sq := squarebox{
		a: image.Pt(200, 50),
		b: image.Pt(400, 50),
		c: image.Pt(200, 250),
		d: image.Pt(400, 250),
	}
	square := drawsquare(sq, gray, 2, color.RGBA{255, 255, 0, 255})
	ii.saveI("line", square)
}
func (ii *imageInfo) grayscaleI(img image.Image) image.Image {
	fmt.Printf("[*] Converting %s to grascale image ...\n", ii.name)
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	ii.saveI("grayscaled", gray)
	return gray
}
func (ii *imageInfo) saveI(name string, img image.Image) {
	out, err := os.Create(fmt.Sprintf("data/%s-%s.gore.%s", name, ii.name, ii.format))
	if err != nil {
		fmt.Printf("image.go:makeItGray:os.Create: image: %s %v\n", name, err)
	}
	defer out.Close()
	fmt.Printf("[*] Saving %s-%s.gore.%s\n", name, ii.name, ii.format)
	switch ii.format {
	case "png":
		png.Encode(out, img)
	case "jpg":
		jpeg.Encode(out, img, nil)
	}
}

func decode(i io.Reader) (image.Image, string) {
	img, f, err := image.Decode(i)
	if err != nil {
		fmt.Printf("error while decoding image: %v\n", err)
		panic("Decode")
	}
	return img, f
}

func (ii *imageInfo) dividI(img image.Image) {
	//divid img to 16x16 images
	bounds := ii.bounds
	w, h := bounds.Max.X, bounds.Max.Y
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("%d, %d", x, y)
		}
	}
}

func hogVect() {
	// http://mccormickml.com/2013/05/07/gradient-vectors/
	// when divided take each block of 16x16 pixel and find where is the highest pixel is.
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
