package core

import (
	"image/color"
	"math"
)

type (
	// RGB color without Alpha
	RGB struct {
		R, G, B float64
	}
	// RGBY Red, Green, Blue and Yellow color
	RGBY struct {
		R, G, B, Y float64
	}
	// XYZ is an additive color space based on how the eye intereprets stimulus from light.
	XYZ struct {
		X, Y, Z float64
	}
	// LAB is CieLab color
	LAB struct {
		L, A, B float64
	}
)

// RGBAtoRGB convert rgba to rgb with a black background.
func RGBAtoRGB(rgba color.RGBA) RGB {
	a := float64(rgba.A)
	if a == 0 {
		return RGB{0, 0, 0}
	}
	return RGB{
		((255-a)*255 + a*float64(rgba.R)) / 255,
		((255-a)*255 + a*float64(rgba.G)) / 255,
		((255-a)*255 + a*float64(rgba.B)) / 255,
	}
}

// RGBtoXYZ convert rgb to xyz.
// source sRGB to XYZ: http://cs.haifa.ac.il/hagit/courses/ist/Lectures/Demos/ColorApplet2/t_convert.html
func RGBtoXYZ(rgb RGB) XYZ {
	r, g, b := rgb.R, rgb.G, rgb.B
	return XYZ{
		float64(r)*0.4124564 + float64(g)*0.3575761 + float64(b)*0.1804375,
		float64(r)*0.2126729 + float64(g)*0.7151521 + float64(b)*0.0721750,
		float64(r)*0.0193339 + float64(g)*0.1191921 + float64(b)*0.9503041,
	}
}

// XYZtoCieLAB convert xyz to Cie*L*a*b
// Reference: https://www.mathworks.com/help/images/ref/xyz2lab.html
// using D65:[0.9504, 1.0000, 1.0888]
func XYZtoCieLAB(xyz XYZ) LAB {
	fx, fy, fz := xyz.X/95.0470, xyz.Y/100.0000, xyz.Y/108.8830
	return LAB{
		L: 116*Ft(fy) - 0.16,
		A: 500 * Ft(fx-fy),
		B: 200 * Ft(fy-fz),
	}
}

// RGBAtoCieLAB convert rgb to CieLAB
// using https://hrcak.srce.hr/file/193994 page 3/8
func RGBAtoCieLAB(rgba color.RGBA) LAB {
	return XYZtoCieLAB(RGBtoXYZ(RGBAtoRGB(rgba)))
}

// Intensity equation to give the intensity of a pixel (RGB)
func Intensity(rgb RGB) float64 {
	return 0.226*rgb.R + 0.7152*rgb.G + 0.0722*rgb.B
}

// Intensityrgb equation to give the intensity of a pixel (RGB)
func Intensityrgb(rgb RGB) (uint8, uint8, uint8) {
	return uint8(0.226 * rgb.R), uint8(0.7152 * rgb.G), uint8(0.0722 * rgb.B)
}

// RGBtoRGBY Convert RGB to Red, Green, Blue and Yellow
func RGBtoRGBY(rgb RGB) RGBY {
	b := rgb.B - (rgb.R+rgb.G)/2
	return RGBY{
		rgbyCondition(rgb.R - (rgb.B+rgb.G)/2),
		rgbyCondition(rgb.G - (rgb.B+rgb.R)/2),
		rgbyCondition(b),
		rgbyCondition(-b - math.Abs(rgb.R-rgb.G)/2),
	}
}
func rgbyCondition(rgby float64) float64 {
	if rgby < 0 {
		return .0
	}
	return rgby
}
