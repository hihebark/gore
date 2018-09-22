package core

import (
	"image/color"
	"math/rand"
	"testing"

	"github.com/hihebark/gore/core"
)

func init() {
	rand.Seed(10)
}

func TestRGBtoXYZ(t *testing.T) {
	rgba := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	t.Logf("%v\n", rgba)
	c := core.RGBtoXYZ(core.RGBAtoRGB(rgba))
	if c.X == 0 && c.Y == 0 && c.Z == 0 {
		t.Errorf("TestRGBAtoXYZ error on converting to xyz nil detect XYZ = %v", c)
	}
	t.Log(c)
}
func TestRGBAtoCieLAB(t *testing.T) {
	rgba := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	t.Logf("%v\n", rgba)
	c := core.XYZtoCieLAB(core.RGBtoXYZ(core.RGBAtoRGB(rgba)))
	if c.L < 0 && c.A < 0 && c.B < 0 {
		t.Errorf("TestRGBAtoXYZ error on converting to xyz nil detect XYZ = %v", c)
	}
	t.Log(c)
}

func TestRGBtoRGBY(t *testing.T) {
	rgba := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	t.Logf("%v\n", rgba)
	c := core.RGBtoRGBY(core.RGBAtoRGB(rgba))
	if c.R < 0 || c.G < 0 || c.B < 0 || c.Y < 0 {
		t.Errorf("TestRGBtoRGBY error value under zero%v", c)
	}
	t.Log(c)
}
