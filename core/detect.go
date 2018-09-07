package core

import (
	"fmt"
	"image"
	"image/color"
	"sort"
)

type Points struct {
	Percent float64
	Rect    image.Rectangle
}
type bypercent []Points

func (b bypercent) Len() int           { return len(b) }
func (b bypercent) Less(i, j int) bool { return b[i].Percent < b[j].Percent }
func (b bypercent) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

//DetectFace if found.
func DetectFace(imgsrc, imgmodel image.Image) []Points {
	fmt.Printf("+ Detecting face(s)...\n")
	bmodel, bsrc := imgmodel.Bounds(), imgsrc.Bounds()
	pts := []Points{}
	total, count := ((bsrc.Max.X - bmodel.Max.X) * (bsrc.Max.Y - bmodel.Max.Y)), 1
	for ys := bsrc.Min.Y; ys+bmodel.Max.Y <= bsrc.Max.Y; ys++ {
		for xs := bsrc.Min.X; xs+bmodel.Max.X <= bsrc.Max.X; xs++ {
			equal := 0
			for ym := bmodel.Min.Y; ym < bmodel.Max.Y; ym++ {
				for xm := bmodel.Min.X; xm < bmodel.Max.X; xm++ {
					if imgsrc.At(xs+xm, ys+ym) == imgmodel.At(xm, ym) && imgsrc.At(xs+xm, ys+ym) == color.White {
						equal++
					}
				}
			}
			percent := float64(equal * 100 / (bmodel.Max.X * bmodel.Max.Y))
			rect := image.Rect(xs, ys, xs+bmodel.Max.X, ys+bmodel.Max.Y)
			pts = append(pts, Points{percent, rect})
			fmt.Printf("- Process %3.0f%%\r", float64(count*100/total))
			count++
		}
	}
	fmt.Println()
	sort.Sort(bypercent(pts))
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].Percent < pts[j].Percent
	})
	return pts[0:10]
}
