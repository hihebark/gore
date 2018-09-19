package core

import "image/color"

//RGB color without Alpha
type RGB struct {
	R, G, B float64
}

//XYZ is an additive color space based on how the eye intereprets stimulus from light.
type XYZ struct {
	X, Y, Z float64
}

//LAB is CieLab color
type LAB struct {
	L, A, B float64
}

// RGBAtoRGB convert rgba to rgb with a black background.
func RGBAtoRGB(rgba color.RGBA) RGB {
	a := float64(rgba.A)
	if a == 0 {
		return RGB{0, 0, 0}
	}
	rgb := RGB{}
	rgb.R = ((255-a)*255 + a*float64(rgba.R)) / 255
	rgb.G = ((255-a)*255 + a*float64(rgba.G)) / 255
	rgb.B = ((255-a)*255 + a*float64(rgba.B)) / 255
	return rgb
}

// RGBtoXYZ convert rgb to xyz.
// source sRGB to XYZ: http://cs.haifa.ac.il/hagit/courses/ist/Lectures/Demos/ColorApplet2/t_convert.html
func RGBtoXYZ(rgb RGB) XYZ {
	xyz := XYZ{}
	r, g, b := rgb.R, rgb.G, rgb.B
	xyz.X = float64(r)*0.4124564 + float64(g)*0.3575761 + float64(b)*0.1804375
	xyz.Y = float64(r)*0.2126729 + float64(g)*0.7151521 + float64(b)*0.0721750
	xyz.Z = float64(r)*0.0193339 + float64(g)*0.1191921 + float64(b)*0.9503041
	return xyz
}

// XYZtoCieLAB convert xyz to Cie*L*a*b
// Reference: https://www.mathworks.com/help/images/ref/xyz2lab.html D65:[0.9504, 1.0000, 1.0888]
func XYZtoCieLAB(xyz XYZ) LAB {
	fx, fy, fz := xyz.X/95.0470, xyz.Y/100.0000, xyz.Y/108.8830
	lab := LAB{}
	lab.L = 116*Ft(fy) - 0.16
	lab.A = 500 * Ft(fx-fy)
	lab.B = 200 * Ft(fy-fz)
	return lab
}

// RGBAtoCieLAB convert rgb to CieLAB
// using https://hrcak.srce.hr/file/193994 page 3/8
func RGBAtoCieLAB(rgba color.RGBA) LAB {
	return XYZtoCieLAB(RGBtoXYZ(RGBAtoRGB(rgba)))
}
