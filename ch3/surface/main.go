// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var (
	InvalidFloatMap map[float64]bool
	sin30, cos30    = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
)

func init() {
	var z float64
	InvalidFloatMap = map[float64]bool{
		1 / z:  true,
		-1 / z: true,
		z / z:  true,
	}
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := cornerV2(i+1, j)
			bx, by, bz := cornerV2(i, j)
			cx, cy, cz := cornerV2(i, j+1)
			dx, dy, dz := cornerV2(i+1, j+1)
			h := (az + bz + cz + dz) / 4
			if !isDataValid(ax, ay, bx, by, cx, cy, dx, dy) {
				// fmt.Printf("i: %v, j: %v\n", i, j)
				// panic("encounter invalid data")
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getColor(h))
		}
	}
	fmt.Println("</svg>")
}

// func main() {
// 	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
// 		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
// 		"width='%d' height='%d'>", width, height)
// 	for i := 0; i < cells; i++ {
// 		for j := 0; j < cells; j++ {
// 			ax, ay := corner(i+1, j)
// 			bx, by := corner(i, j)
// 			cx, cy := corner(i, j+1)
// 			dx, dy := corner(i+1, j+1)
// 			if !isDataValid(ax, ay, bx, by, cx, cy, dx, dy) {
// 				// fmt.Printf("i: %v, j: %v\n", i, j)
// 				// panic("encounter invalid data")
// 				continue
// 			}
// 			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='blue'/>\n",
// 				ax, ay, bx, by, cx, cy, dx, dy)
// 		}
// 	}
// 	fmt.Println("</svg>")
// }

// func corner(i, j int) (float64, float64) {
// 	// Find point (x,y) at corner of cell (i,j).
// 	x := xyrange * (float64(i)/cells - 0.5)
// 	y := xyrange * (float64(j)/cells - 0.5)

// 	// Compute surface height z.
// 	z := f(x, y)

// 	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
// 	sx := width/2 + (x-y)*cos30*xyscale
// 	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
// 	return sx, sy
// }

func cornerV2(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func isDataValid(data ...float64) bool {
	for _, item := range data {
		if math.IsNaN(item) {
			return false
		}
		if math.IsInf(item, 0) {
			return false
		}
	}
	return true
}

func normalize(x float64) float64 {
	max := 1.0
	min := -0.5
	if x > max {
		return 1.0
	}
	if x < min {
		return 0.0
	}
	return (x - min) / (max - min)
}

func getColor(h float64) string {
	r := normalize(h) * 255.0
	b := (1 - normalize(h)) * 255.0

	return fmt.Sprintf("#%02X00%02X", int(r), int(b))
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// a saddle
// func f(x, y float64) float64 {
// 	a, b := 20.0, 10.0
// 	// r := math.Hypot(x, y) // distance from (0,0)
// 	return x*x/a/a - y*y/b/b
// }

// moguls
// func f(x, y float64) float64 {
// 	r := math.Sin(x/3) + math.Cos(y/3)
// 	return r
// }
