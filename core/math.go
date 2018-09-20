package core

import (
	"math"
)

// Constant
const (
	FULLCIRCLE float64 = 360
	HALFCIRCLE float64 = 180
	K          float64 = 8
)

// Magnitude calculate the magnitude of two points
//		f(x, y) = sqrt(c^2 + y^2)
func Magnitude(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

// OrientationXY calculate orientation of two points
//		f(x, y) = atan2(x, y) * 180 / 3.14 % 360
func OrientationXY(x, y float64) float64 {
	return math.Mod((math.Atan2(x, y) * HALFCIRCLE / math.Pi), FULLCIRCLE)
}

func xFromAngle(x, length int, angle float64) float64 {
	return math.Round(float64(x) + (float64(length) * math.Cos(angle)))
}
func yFromAngle(y, length int, angle float64) float64 {
	return math.Round(float64(y) + (float64(length) * math.Sin(angle)))
}

// Gaussian formula.
//		G(x, y) = (1/2PI*sigma^2)(exp(-x^2+y^2/2sigma^2))
func Gaussian(x, y int, sigma float64) float64 {
	return math.Exp(float64(-(x*x+y*y))/(2*sigma*sigma)) / (2 * math.Pi * sigma * sigma)
}

//Ft function
func Ft(t float64) float64 {
	if t >= 0.008856 {
		//return math.Cbrt(t)
		return math.Pow(t, 1/3)
	}
	return 7.787037*t + 0.13793
}

// Intensity equation to give the intensity of a pixel (RGB)
func Intensity(rgb RGB) float64 {
	return 0.226*rgb.R + 0.7152*rgb.G + 0.0722*rgb.B
}

// Orientation not now
func Orientation() {

}

// Gabor filter.
//		G(θ, λ, ϕ, σ, γ, x,y) = exp(x'^2+(y^2*y'^2)/2σ^2) * exp(i(2PI*x'/ γ + ϕ))
//		where
//			x' = x*cosθ + y*sinθ
//			y' = x*sinθ + y*cosθ
//			σ = sigma of gaussian envelope
//			γ = gamma spatial aspect ratio
//			ϕ = phi phase offset
//			θ = theta angle
//			λ = lambda
func Gabor(x, y, lambda int) []float64 {
	gamma := 1.0
	sigma := math.Pi
	phi := .0
	gabors := make([]float64, 8)
	//thetas := [8]string{"0", "pi/8", "2pi/8", "3pi/8", "4pi/8", "5pi/8", "6pi/8", "7pi/8"}
	for i := 0; i < int(K); i++ {
		theta := float64(i) * math.Pi / K
		xx := (float64(x)*math.Cos(theta) + float64(y)*math.Sin(theta))
		yy := (float64(x)*math.Sin(theta) + float64(y)*math.Cos(theta))
		yy = math.Pow(yy, 2)
		gabor := math.Exp(math.Pow(xx, 2) + (yy*math.Pow(gamma, 2))/math.Pow(2*sigma, 2))
		gabor *= math.Exp(float64(i) * ((2.0 * math.Pi * xx / float64(lambda)) + phi))
		gabors[i] = gabor
	}
	return gabors
}
