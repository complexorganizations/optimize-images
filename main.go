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
	imageLocation = "example.png"
	if len(os.Args) > 1 {
		imageLocation = os.Args[1]
	}
}

func main() {
	if fileExists(imageLocation) {
		imageformatChoice()
	} else if folderExists(imageLocation) {
		formatAllImages()
	}
}

func formatAllImages() {
	filepath.Walk(imageLocation, func(path string, info os.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".jpeg", ".jpg":
			jpegImage(path)
		case ".png":
			pngImage(path)
		case ".gif":
			gifImage(path)
		default:
			fmt.Println("Error:", imageLocation)
		}
		return nil
	})
}

func imageformatChoice() {
	switch filepath.Ext(imageLocation) {
	case ".jpeg", ".jpg":
		jpegImage(imageLocation)
	case ".png":
		pngImage(imageLocation)
	case ".gif":
		gifImage(imageLocation)
	default:
		fmt.Println("Error:", imageLocation)
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
	fmt.Println("Enhancing:", imageLocation)
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
	fmt.Println("Enhancing:", imageLocation)
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
	fmt.Println("Enhancing:", imageLocation)
}

func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
