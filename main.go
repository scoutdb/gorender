package main

import (
	"fmt"
	"io/ioutil"
)

const (
	image_heigth int = 256
	image_width  int = 256
)

type output struct {
	photo string
}

// https://raytracing.github.io/books/RayTracingInOneWeekend.html#outputanimage
func main() {

	output := &output{}
	output.photo = fmt.Sprintf("P3\n%v %v \n255\n", image_width, image_heigth)
	for i := image_heigth - 1; i >= 0; i-- {
		for j := 0; j < image_width; j++ {
			// fmt.Println(i)
			r := float32(j) / float32(image_width-1)
			g := float32(i) / float32(image_heigth-1)
			b := float32(0.25)

			ir := float32(255.999) * r
			ig := float32(255.999) * g
			ib := float32(255.999) * b
			//once the calculation is done i need to convert back to int
			br := int(ir)
			bg := int(ig)
			bb := int(ib)

			output.photo += fmt.Sprintf("%v %v %v\n", br, bg, bb)
		}
	}
	err := ioutil.WriteFile("image.ppm", []byte(output.photo), 0755)
	if err != nil {
		fmt.Println(err)
	}
}
