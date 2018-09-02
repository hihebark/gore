package core

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"sync"

	"golang.org/x/image/draw"
)

//ImageInfo image information.
type ImageInfo struct {
	Wg sync.WaitGroup
	sync.RWMutex
	Format   string
	Name     string
	Bounds   image.Rectangle
	Scalsize int
	Cellsize int
}

//NewImageInfo return ImageInfo struct.
func NewImageInfo(f, n string, b image.Rectangle, s, c int) *ImageInfo {
	return &ImageInfo{
		Wg:       sync.WaitGroup{},
		Format:   f,
		Name:     n,
		Bounds:   b,
		Scalsize: s,
		Cellsize: c,
	}
}

//Gray scale image
func (i *ImageInfo) Grayscale(imgsrc image.Image) image.Image {
	fmt.Printf("+ Grascaling image ...\n")
	if imgsrc.ColorModel() == color.GrayModel {
		return imgsrc
	}
	bounds := imgsrc.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			gray.Set(x, y, color.GrayModel.Convert(imgsrc.At(x, y)))
		}
	}
	return gray
}

//Save save image into directory
func (i *ImageInfo) Save(name string, imgsrc image.Image) {
	out, err := os.Create(fmt.Sprintf("data/%s-%s.gore.%s", name, i.Name, i.Format))
	if err != nil {
		fmt.Printf("image.go:makeItGray:os.Create: image: %s %v\n", name, err)
	}
	defer out.Close()
	fmt.Printf("+ Saving %s-%s.gore.%s\n", name, i.Name, i.Format)
	switch i.Format {
	case "png":
		png.Encode(out, imgsrc)
	case "jpg", "jpeg":
		jpeg.Encode(out, imgsrc, nil)
	}
}

//Scale reduce image into i.scalsize defind in ImageInfo.
func (i *ImageInfo) Scale(imgsrc image.Image) image.Image {
	fmt.Printf("+ Scale image into %d\n", i.Scalsize)
	bound := imgsrc.Bounds()
	rect := image.Rect(0, 0, int(bound.Max.X/i.Scalsize), int(bound.Max.Y/i.Scalsize))
	dstimg := image.NewRGBA(rect)
	draw.ApproxBiLinear.Scale(dstimg, rect, imgsrc, imgsrc.Bounds(), draw.Over, nil)
	return dstimg
}

func decode(i io.Reader) (image.Image, string) {
	img, f, err := image.Decode(i)
	if err != nil {
		fmt.Printf("error while decoding image: %v\n", err)
		panic("Decode")
	}
	return img, f
}

//Divid split rectangle into s*s rectangles.
func Divid(bounds image.Rectangle, s int) []image.Rectangle {
	w, h, c := bounds.Max.X, bounds.Max.Y, 0
	cells := make([]image.Rectangle, int(w/s*h/s))
	for y := 16; y < h; y += s {
		for x := 16; x < w; x += s {
			v, z := x, y
			cells[c] = image.Rect(v-s, z-s, x, y)
			c++
		}
	}
	return cells
}

/************************************************************************
type pixel struct {
	R, G, B, A int
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
************************************************************************/
