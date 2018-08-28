package core

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"
	"sync"

	"golang.org/x/image/draw"
)

type pixel struct {
	R, G, B, A int
}
type ImageInfo struct {
	wg sync.WaitGroup
	sync.RWMutex
	format   string
	name     string
	bounds   image.Rectangle
	scalsize int
	cellsize int
}

func NewImageInfo(f, n string, b image.Rectangle, s, c int) *ImageInfo {
	return &ImageInfo{
		wg:       sync.WaitGroup{},
		format:   f,
		name:     n,
		bounds:   b,
		scalsize: s,
		cellsize: c,
	}
}

//Gray scale image
func (i *ImageInfo) Grayscale(img image.Image) image.Image {
	fmt.Printf("+ Grascaling image ...\n")
	if img.ColorModel() == color.GrayModel {
		return img
	}
	bounds := img.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return gray
}

//Save save image into directory
func (i *ImageInfo) Save(name string, img image.Image) {
	out, err := os.Create(fmt.Sprintf("data/%s-%s.gore.%s", name, i.name, i.format))
	if err != nil {
		fmt.Printf("image.go:makeItGray:os.Create: image: %s %v\n", name, err)
	}
	defer out.Close()
	fmt.Printf("+ Saving %s-%s.gore.%s\n", name, i.name, i.format)
	switch i.format {
	case "png":
		png.Encode(out, img)
	case "jpg", "jpeg":
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

//HogVect hog implementation.
func (i *ImageInfo) HogVect(img image.Image) image.Image {
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)
	cells := i.divid(img)
	midcell := image.Pt(int(i.cellsize/2), int(i.cellsize/2))
	c := color.RGBA{0xff, 0xff, 0xff, 0xbb}
	fmt.Printf("+ There is %d cells\n", len(cells)-1)
	for k, cell := range cells {
		i.wg.Add(1)
		fmt.Printf("- Processing with %d cell\r", k)
		imgcell := image.NewRGBA(cell)
		for y := cell.Min.Y; y < cell.Max.Y; y++ {
			for x := cell.Min.X; x < cell.Max.X; x++ {
				yd := math.Abs(float64(img.At(x, y-1).(color.Gray).Y - img.At(x, y+1).(color.Gray).Y))
				xd := math.Abs(float64(img.At(x-1, y).(color.Gray).Y - img.At(x+1, y).(color.Gray).Y))
				magnitude, orientation := gradientVector(xd, yd)
				imgcell = drawLine(cell.Sub(midcell).Max, orientation, magnitude, imgcell, c)
			}

		}
		draw.Draw(dst, imgcell.Bounds(), imgcell, cell.Min, draw.Over)
		i.wg.Done()
	}
	fmt.Print("\n")
	return dst
}
func (i *ImageInfo) divid(img image.Image) []image.Rectangle {
	//divid img to 16x16 cells
	bounds := img.Bounds()
	w, h, c := bounds.Max.X, bounds.Max.Y, 0
	s := i.cellsize
	cells := make([]image.Rectangle, int(w*h/(s*s)+1)) // TODO not sure if it's correcte to verify later.
	for y := 16; y < h; y += s {
		for x := 16; x < w; x += s {
			v, z := x, y
			cells[c] = image.Rect(v-s, z-s, x, y)
			c++
		}
	}
	return cells
}
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
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) pixel {
	return pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}
