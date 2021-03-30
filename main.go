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

var systemPath string
var files []string

func init() {
	systemPath = "/"
	if len(os.Args) > 1 {
		systemPath = os.Args[1]
	}
}

func main() {
	formatAllImages()
}

func formatAllImages() {
	if systemPath == "/" {
		filepath.Walk(systemPath, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
	}
	fmt.Println(files)
	imageformatChoice()
}

func imageformatChoice() {
	if systemPath == "/" {
		systemPath = files
	}
	switch filepath.Ext(systemPath) {
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
	file, err := os.Open(systemPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(systemPath)
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
	file, err := os.Open(systemPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(systemPath)
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
	file, err := os.Open(systemPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	outfile, err := os.Create(systemPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()
	gif.Encode(outfile, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}
