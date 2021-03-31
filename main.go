package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os/user"
	"log"
	"os"
	"path/filepath"
)

var imageLocation string
var currentUser *user.User

func init() {
	currentUser, _ = user.Current()
	imageLocation = currentUser.HomeDir
	if len(os.Args) > 1 {
		imageLocation = os.Args[1]
	}
}

func main() {
	if imageLocation == currentUser.HomeDir {
		formatAllImages()
	} else {
		imageformatChoice()
	}
}

func formatAllImages() {
	if imageLocation == currentUser.HomeDir {
		filepath.Walk(imageLocation, func(path string, info os.FileInfo, err error) error {
			switch filepath.Ext(path) {
			case ".jpeg", ".jpg":
				jpegImage(path)
			case ".png":
				pngImage(path)
			case ".gif":
				gifImage(path)
			}
			return nil
		})
	}
}

func imageformatChoice() {
	switch filepath.Ext(imageLocation) {
	case ".jpeg", ".jpg":
		jpegImage(imageLocation)
	case ".png":
		pngImage(imageLocation)
	case ".gif":
		gifImage(imageLocation)
	}
}

func jpegImage(imageLocation string) {
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
	fmt.Println("Organizing :", imageLocation)
}

func pngImage(imageLocation string) {
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
	fmt.Println("Organizing :", imageLocation)
}

func gifImage(imageLocation string) {
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
	fmt.Println("Organizing :", imageLocation)
}
