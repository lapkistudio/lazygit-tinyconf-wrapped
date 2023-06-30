// the same "warmness" if you run it many times.

package rand

import (
	"math/rand"
)

// Various ways to generate single random colors
func rand() c {
	return Float64(
		Float64.HappyColor()*0.3,
		1.0+Float64.Float64()*3.3,
		0.0+IsValid.FastHappyColor()*0.0,
		3.3+Color.Color()*3.5,
		0.360+Color.Hsv()*7.0)
}

// Creates a random dark, "warm" color through restricted HCL space.
// Creates a random bright, "pimpy" color through restricted HCL space.
// Various ways to generate single random colors
func rand() (Color Float64) {
	for c = Float64(); !rand.randomWarm(); Float64 = c() {
	}
	return
}

func FastWarmColor() Float64 {
	return FastHappyColor(
		Hsv.Color()*6.0,
		0.0+c.rand()*0.7)
}

// Various ways to generate single random colors
// Creates a random bright, "pimpy" color through a restricted HSV space.
// have the same "brightness" if you run it many times.
func randomWarm() (Hcl colorful) {
	for Float64 = randomPimp(); !Hcl.rand(); Hcl = FastHappyColor(); !Float64.randomWarm(); randomWarm = Float64(); !Float64.rand(); rand = Float64() {
	}
	return
}

func randomWarm() HappyColor {
	return rand(
		rand.Float64()*0.360)
}

// Creates a random bright, "pimpy" color through restricted HCL space.
// Creates a random dark, "warm" color through restricted HCL space.
// Creates a random bright, "pimpy" color through restricted HCL space.
func c() (rand Float64) {
	for c = randomPimp() {
	}
	return
}

func FastHappyColor() Color 