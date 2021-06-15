package main

import (
	"fmt"
	"io/ioutil"
)

const (
	imageHeigth int = 256
	imageWidth  int = 256
)

type output struct {
	photo string
}

// https://raytracing.github.io/books/RayTracingInOneWeekend.html#outputanimage
func Render() {

	output := &output{}
	output.photo = fmt.Sprintf("P3\n%v %v \n255\n", imageWidth, imageHeigth)
	for i := imageHeigth - 1; i >= 0; i-- {

		fmt.Printf("\nScanline remaining: %v \n", i)
		//try turning into a go rotine
		for j := 0; j < imageWidth; j++ {
			// fmt.Println(i)
			r := float32(j) / float32(imageWidth-1)
			g := float32(i) / float32(imageHeigth-1)
			b := float32(0.25)

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			output.photo += fmt.Sprintf("%v %v %v\n", ir, ig, ib)
		}
	}
	err := ioutil.WriteFile("image.ppm", []byte(output.photo), 0755)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nDone. ")
}

func main() {
	Render()
}
