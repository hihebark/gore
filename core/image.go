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
	"strings"
)

type pixel struct {
	R int
	G int
	B int
	A int
}
type imageInfo struct {
	format   string
	name     string
	bounds   image.Rectangle
	sizescal int
}

func newImageInfo(f, n string, b image.Rectangle, s int) *imageInfo {
	return &imageInfo{
		format:   f,
		name:     n,
		bounds:   b,
		sizescal: s,
	}
}

//Start detecting face in image given.
func Start(path string) {
	img, err := os.Open(path)
	defer img.Close()
	if err != nil {
		fmt.Printf("image:start:os.Open path:%s\n", path)
	}

	info, _ := img.Stat()
	name := strings.Split(info.Name(), ".")[0]
	imgdec, form := decode(img)
	imginf := newImageInfo(form, name, imgdec.Bounds(), 2)
	gray := grayscaleI(imgdec)
	imginf.saveI("grayscaled", gray)
	imginf.saveI("lines", hogVect(scaleImage(gray, 2)))
	/*************************************************************
	line := drawLine(image.Pt(5, 5), 44.88743476267866, 10, gray)
	imginf.saveI("line", line)
	imginf.saveI("sq",
		drawsquare(gray,
					image.Rect(200, 50, 400, 250),
					2,
					color.RGBA{255, 255, 0, 255}))
	scaledimg := scaleImage(gray, 2)
	nimg := hogVect(scaledimg)
	imginf.saveI("hog", nimg)
	hogVect(gray)
	**************************************************************/
}
func grayscaleI(img image.Image) image.Image {
	fmt.Printf("[*] Grascaling image ...\n")
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
func (i *imageInfo) saveI(name string, img image.Image) {
	out, err := os.Create(fmt.Sprintf("data/%s-%s.gore.%s", name, i.name, i.format))
	if err != nil {
		fmt.Printf("image.go:makeItGray:os.Create: image: %s %v\n", name, err)
	}
	defer out.Close()
	fmt.Printf("[*] Saving %s-%s.gore.%s\n", name, i.name, i.format)
	switch i.format {
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

func hogVect(img image.Image) image.Image {
	nimg := newImage(img.Bounds(), color.RGBA{0, 0, 0, 255})
	cells := dividI(img, 16)
	fmt.Printf("[*] There is %d cells\n", len(cells))
	for k, cell := range cells {
		fmt.Printf("[!] Processing with %d cell\r", k)
		imgcell := newImage(cell, color.RGBA{0, 0, 0, 255})
		for y := cell.Min.Y; y < cell.Max.Y; y++ {
			for x := cell.Min.X; x < cell.Max.X; x++ {
				yd := math.Abs(float64(img.At(x, y-1).(color.Gray).Y - img.At(x, y+1).(color.Gray).Y))
				xd := math.Abs(float64(img.At(x-1, y).(color.Gray).Y - img.At(x+1, y).(color.Gray).Y))
				magnitude, orientation := gradientVector(xd, yd)
				imgcell = drawLine(image.Pt(cell.Min.X, cell.Max.Y), orientation, magnitude, imgcell)
				nimg = squaretoImage(imgcell, nimg, imgcell.Bounds(), image.Pt(cell.Min.X, cell.Min.Y))
				//fmt.Printf("mag:%v ori:%v ", magnitude, orientation)
			}
		}
		//TODO array
	}
	fmt.Println("")
	return nimg
}
func dividI(img image.Image, s int) []image.Rectangle {
	//divid img to 16x16 cells
	bounds := img.Bounds()
	w, h, i := bounds.Max.X, bounds.Max.Y, 0
	cells := make([]image.Rectangle, int(w*h/(s*s))+1) // TODO not sure if it's correcte to verify later.
	for y := 16; y < h; y += s {
		for x := 16; x < w; x += s {
			v, z := x, y
			cells[i] = image.Rect(v-s, z-s, x, y)
			i++
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
