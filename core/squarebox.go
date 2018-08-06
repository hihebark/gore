package core

import (
	"image"
	"image/color"
	//	"golang.org/x/image/draw"
)

type squarebox struct {
	a, b, c, d image.Point
}

//DrawLine draw line
func drawsquare(sq squarebox, img image.Image, density int, c color.Color) image.Image {
	bounds := img.Bounds()
	nimg := image.NewRGBA(bounds)
	w, h := bounds.Max.X, bounds.Max.Y
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			nimg.Set(x, y, img.At(x, y))
			if (y == sq.a.Y && (x >= sq.a.X && x <= sq.b.X)) ||
				(y == sq.c.Y && (x >= sq.c.X && x <= sq.d.X)) ||
				(x == sq.a.X && (y >= sq.a.Y && y <= sq.c.Y)) ||
				(x == sq.b.X && (y >= sq.b.Y && y <= sq.d.Y)) {

				for i := density; i >= 0; i-- {
					nimg.Set(x-i, y-i, c)
					nimg.Set(x+i, y+i, c)
				}
			}
		}
	}
	return nimg
}

/******************************************************************************************
func DrawSB(p image.Point, img image.Image) image.Image {
	s, err := os.Open("data/squarebox.png")
	if err != nil {
		fmt.Printf("error:DrawSB: %v", err)
	}
	square, _ := decode(s)
	dst := image.NewRGBA(img.Bounds())
	//draw.Copy(dst, image.Pt(100, 100), square, square.Bounds(), draw.Src, nil)
	draw.Draw(dst, img.Bounds(), square, square.Bounds().Min, draw.Over)
	//draw.ApproxBiLinear.Scale(dst, dst.Bounds(), square, square.Bounds(), draw.Over, nil)
	//draw.DrawMask(dst, img.Bounds(), square, image.ZP, square, image.ZP, draw.Over)
	return img
}
********************************************************************************************/
