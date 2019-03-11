package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
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
	fmt.Println("Sorting colors")
	fmt.Println("This may take a few minutes...")

	// LONGEST PART OF PROCESS
	clrs = quicksort(clrs)

	fmt.Println("Creating file...")
	var lines []string

	// SECOND LONGEST PART OF PROCESS
	for _, c := range clrs {
		//fmt.Println(c.r, c.g, c.b)
		fmt.Printf("\033\r" + toString(int(c.r)) + " " + toString(int(c.g)) + " " + toString(int(c.b)) + " \r")
		lines = append(lines, toString(int(c.r))+" "+toString(int(c.g))+" "+toString(int(c.b)))
	}
	fmt.Printf("\r\f")
	writeDataFile("data.txt", lines)
	fmt.Println("Done")
	//open.StartWith("data.txt", "notepad")
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

func writeDataFile(filepath string, lines []string) {
	data := strings.Join(lines, "\n")
	err := ioutil.WriteFile(filepath, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func toString(in int) string {
	return fmt.Sprintf("%d", in)
}
