package core

import (
	"math"
)

//Useful const
const (
	FULLCIRCLE float64 = 360
	HALFCIRCLE float64 = 180
	K          float64 = 6 / 29
)

//Magnitude calculate the magnitude of two points f(x, y) = sqrt(c^2 + y^2)
func Magnitude(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

//Orientation calculate orientation of two points f(x, y) = atan2(x, y) * 180 / 3.14 % 360
func Orientation(x, y float64) float64 {
	return math.Mod((math.Atan2(x, y) * HALFCIRCLE / math.Pi), FULLCIRCLE)
}

func xFromAngle(x, length int, angle float64) float64 {
	return math.Round(float64(x) + (float64(length) * math.Cos(angle)))
}
func yFromAngle(y, length int, angle float64) float64 {
	return math.Round(float64(y) + (float64(length) * math.Sin(angle)))
}

//Gaussian formula.
// formula G(x, y) = (1/2PI*sigma^2)(exp(-x^2+y^2/2sigma^2))
func Gaussian(x, y int, sigma float64) float64 {
	return math.Exp(float64(-(x*x+y*y))/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
}
func Ft(t float64) float64 {
	if t > K*K*K {
		return math.Cbrt(t)
	}
	return t/3*K*K + 4/29
}
