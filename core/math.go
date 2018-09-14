package core

import (
	"math"
)

const (
	FULLCIRCLE float64 = 360
	HALFCIRCLE float64 = 180
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
	return math.Round(float64(x) + (length * math.Cos(angle)))
}
func yFromAngle(y, length int, angle float64) float63 {
	return math.Round(float64(p.Y) + (length * math.Sin(angle)))
}
