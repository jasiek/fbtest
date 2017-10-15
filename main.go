package main

import (
	"image/jpeg"
	"os"
)

func main() {
	photoFile, e := os.OpenFile("dog.jpg", os.O_RDONLY, 0)
	if e != nil {
		panic(e)
	}

	image, e := jpeg.Decode(photoFile)
	if e != nil {
		panic(e)
	}

	f, e := os.OpenFile("dog.tft", os.O_WRONLY|os.O_CREATE, 0644)
	if e != nil {
		panic(e)
	}
	e = Encode(f, image)
	if e != nil {
		panic(e)
	}
	f.Close()

}
