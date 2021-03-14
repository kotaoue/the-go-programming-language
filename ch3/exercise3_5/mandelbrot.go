// $ go run ch3/exercise3_5/mandelbrot.go > ch3/exercise3_5/result.png;open -a '/Applications/Google Chrome.app' ch3/exercise3_5/result.png
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin   = -2
		ymin   = -2
		xmax   = 2
		ymax   = 2
		width  = 640
		height = 640
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)

			img.Set(px, py, mandelbrot(x, y, z))
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(x, y float64, z complex128) color.RGBA {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := (255 - contrast*n)
			g := (255 - contrast*n)
			b := (255 - contrast*n)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}