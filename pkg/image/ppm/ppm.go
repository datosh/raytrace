package ppm

import (
	"fmt"
	"image"
	"io"
)

func imageSize(m image.Image) (int, int) {
	imageSize := m.Bounds().Size()
	width := imageSize.X
	height := imageSize.Y
	return width, height
}

func readRGBAt(m image.Image, x, y int) (r, g, b uint8) {
	br, bg, bb, _ := m.At(x, y).RGBA()
	return uint8(br), uint8(bg), uint8(bb)
}

func writeHeader(w io.Writer, width, height int) {
	fmt.Fprintf(w, "P3\n%d %d\n255\n", width, height)
}

func writePixel(w io.Writer, r, g, b uint8) {
	fmt.Fprintf(w, "%d %d %d\n", r, g, b)
}

func Encode(w io.Writer, m image.Image) error {
	width, height := imageSize(m)
	writeHeader(w, width, height)

	for x := height; x > 0; x-- {
		for y := width; y > 0; y-- {
			r, g, b := readRGBAt(m, x, y)
			writePixel(w, r, g, b)
		}
	}
	return nil
}
