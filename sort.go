package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

type clr struct {
	r, g, b float64
}

var iterations int

func main() {
	iterations = 0
	var clrs []clr
	fmt.Println("Creating colors...")
	// Runs in O(scary) time. Seems fast enough in production
	for r := 0.0; r < 256; r++ {
		for g := 0.0; g < 256; g++ {
			for b := 0.0; b < 256; b++ {
				clrs = append(clrs, clr{r, g, b})
			}
		}
	}
	fmt.Println("Sorting colors...")
	fmt.Println("This may take a few minutes")

	// LONGEST PART OF PROCESS
	clrs = quicksort(clrs)

	// make png
	fmt.Println("Creating png...")
	width := 8192
	height := 1024
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			setColor := clrs[x*len(clrs)/width] //color.RGBA{100, 200, 200, 0xff}
			setColorRGBA := color.RGBA{uint8(setColor.r), uint8(setColor.g), uint8(setColor.b), 0xff}
			img.Set(x, y, setColorRGBA)
		}
	}
	// export png
	f, _ := os.Create("image.png")
	png.Encode(f, img)
	f.Close()
	fmt.Println("Done")
}

func (c clr) getFactor() float64 { // returns luminance of a color
	return c.r*0.5 + c.g*0.7 + c.b*0.2
}

func quicksort(in []clr) []clr {
	if len(in) < 2 {
		return in
	}
	left, right := 0, len(in)-1
	pivot := rand.Int() % len(in)
	in[pivot], in[right] = in[right], in[pivot]
	for i := range in {
		if in[i].getFactor() < in[right].getFactor() {
			in[left], in[i] = in[i], in[left]
			left++
		}
	}
	in[left], in[right] = in[right], in[left]
	iterations++

	quicksort(in[:left])
	quicksort(in[left+1:])
	return in
}
