package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	path := "image.jpeg"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	jpeg.Encode(outfile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
