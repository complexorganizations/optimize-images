package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

var imageLocation string

func init() {
	if len(os.Args) > 1 {
		imageLocation = os.Args[1]
	} else {
		log.Fatal("Error: No argument passed.")
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
		if fileExists(path) {
			switch filepath.Ext(path) {
			case ".jpeg", ".jpg":
				jpegImage(path)
			case ".png":
				pngImage(path)
			default:
				log.Println("Warning:", imageLocation)
			}
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
	default:
		log.Println("Warning:", imageLocation)
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
	log.Println("Enhancing:", imageLocation)
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
	log.Println("Enhancing:", imageLocation)
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
