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

//Grayscale gray scale image
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

//Divid split rectangle into s*s cell.
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

//Blur image.
func Blur(imgsrc image.Image, radius float64) image.Image {
	maxY, maxX := imgsrc.Bounds().Max.Y, imgsrc.Bounds().Max.X
	imgdst := image.NewRGBA(imgsrc.Bounds())
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			var r, g, b, a uint32 = 0, 0, 0, 0
			var count uint32

			for ky := -radius; ky < radius; ky++ {
				for kx := -radius; kx <= radius; kx++ {
					kr, kg, kb, ka := imgsrc.At(x+int(kx), y+int(ky)).RGBA()
					r += kr
					g += kg
					b += kb
					a += ka
					count++
				}
			}
			c := color.RGBA{uint8(r/count) + 1, uint8(g/count) + 1, uint8(b/count) + 1, uint8(a / count)}
			imgdst.Set(x, y, c)
		}
	}
	return imgdst
}

//GaussianBlur blur image with gauss formula.
func GaussianBlur(imgsrc image.Image, kernel, radius int) image.Image {
	bounds := imgsrc.Bounds()
	maxX, maxY := bounds.Max.X, bounds.Max.Y
	imgdst := image.NewRGBA64(bounds)
	l := maxY * maxX
	kernels := gaussianMap(kernel, float64(radius))
	fmt.Printf("+ There is %d cells\n", l)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			var r, g, b, a uint16

			for ky := -kernel; ky < kernel; ky++ {
				for kx := -kernel; kx < kernel; kx++ {
					kr, kg, kb, ka := imgsrc.At(x+kx, y+ky).RGBA()
					r += uint16(float64(kr) * kernels[kernel+kx][kernel+ky])
					g += uint16(float64(kg) * kernels[kernel+kx][kernel+ky])
					b += uint16(float64(kb) * kernels[kernel+kx][kernel+ky])
					a += uint16(float64(ka) * kernels[kernel+kx][kernel+ky])
				}
			}
			imgdst.SetRGBA64(x, y, color.RGBA64{r, g, b, a})
			fmt.Printf("- Processing with %5d cell\r", (maxX*maxY)-l)
			l--
		}
	}
	fmt.Printf("\n")
	return imgdst
}

func gaussianMap(ks int, sigma float64) [][]float64 {
	var sum float64
	l := ks*2 + 1
	kernel := make([][]float64, l)
	for i := 0; i < l; i++ {
		row := make([]float64, l)
		for j := 0; j < l; j++ {
			g := Gaussian(i, j, sigma)
			row[j] = g
			sum += g
		}
		kernel[i] = row
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			kernel[i][j] /= sum
		}
	}
	return kernel
}
