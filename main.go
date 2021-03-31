package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var imageLocation string

func init() {
	imageLocation = "/"
	if len(os.Args) > 1 {
		imageLocation = os.Args[1]
	}
}

func main() {
	imageformatChoice()
}

func formatAllImages() {
	if imageLocation == "/" {
		filepath.Walk(imageLocation, func(path string, info os.FileInfo, err error) error {
			//
			return nil
		})
	}
}

func imageformatChoice() {
	switch filepath.Ext(imageLocation) {
	case ".jpeg", ".jpg":
		jpegImage()
	case ".png":
		pngImage()
	case ".gif":
		gifImage()
	default:
		fmt.Println("Error: Image format not supported")
	}
}

func jpegImage() {
	file, err := os.Open(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	jpeg.Encode(outfile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pngImage() {
	file, err := os.Open(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	png.Encode(outfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func gifImage() {
	file, err := os.Open(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(imageLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	gif.Encode(outfile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
