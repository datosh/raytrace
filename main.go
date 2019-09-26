package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	ppm "github.com/datosh/raytrace/pkg/image/ppm"
)

func openExclusive(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC, 0)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}

func storeAsPPM(img image.Image, filename string) error {
	file := openExclusive(filename)
	defer file.Close()

	ppm.Encode(file, img)
	return nil
}

func storeAsPNG(img image.Image, filename string) error {
	file := openExclusive(filename)
	defer file.Close()

	png.Encode(file, img)
	return nil
}

func main() {
	img := image.NewRGBA(image.Rect(-100, -50, 100, 50))
	fmt.Printf("Image has size of %dx%d", img.Rect.Dx(), img.Rect.Dy())
	for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {

			dr := float64(x+img.Rect.Min.X) / float64(img.Rect.Dx())
			dg := float64(y+img.Rect.Min.Y) / float64(img.Rect.Dy())
			img.SetRGBA(x, y, color.RGBA{
				R: 255 - uint8(dr*255.99),
				G: uint8(dg * 255.99),
				B: 0,
				A: 255,
			})
		}
	}

	storeAsPNG(img, "image.png")
}
