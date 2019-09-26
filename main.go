package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	ray "github.com/datosh/raytrace/pkg/raytracing"

	vector "github.com/datosh/raytrace/pkg/vector"

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

func cameraDirection(u, v float64) vector.Vec3 {
	lowerLeft := vector.NewVec3(-2.0, -1.0, -1.0)
	horizontal := vector.NewVec3(4.0, 0.0, 0.0)
	vertical := vector.NewVec3(0.0, 2.0, 0.0)

	vv := vector.Scale(vertical, v)
	uv := vector.Scale(horizontal, u)
	return vector.Add(lowerLeft, vector.Add(vv, uv))
}

func vecToColor(v vector.Vec3) color.RGBA {
	return color.RGBA{
		R: uint8(v.X * 255.0),
		G: uint8(v.Y * 255.0),
		B: uint8(v.Z * 255.0),
		A: 255,
	}
}

// Implements a lerp (lineral interpolation)
// Depending on t, which is up/down-ness of direction from camera to image
// The color goes from white to blue
// lerp(t) = (1-t)*from + t*to
func rayToColor(r ray.Ray) color.RGBA {
	if hitSphere(vector.NewVec3(0.0, 0.0, -1.0), 0.8, r) {
		return vecToColor(vector.NewVec3(1.0, 0.0, 0.0))
	}
	unitDir := vector.Unit(r.Direction())
	t := 0.5 * (unitDir.Y + 1.0)
	a := vector.Scale(vector.NewVec3(1.0, 1.0, 1.0), (1.0 - t))
	b := vector.Scale(vector.NewVec3(0.2, 0.5, 1.0), t)
	return vecToColor(vector.Add(a, b))
}

func hitSphere(center vector.Vec3, radius float64, r ray.Ray) bool {
	oc := vector.Sub(r.Origin(), center)
	a := vector.Dot(r.Direction(), r.Direction())
	b := 2.0 * vector.Dot(oc, r.Direction())
	c := vector.Dot(oc, oc) - radius*radius
	disc := b*b - 4*a*c
	return disc > 0
}

func main() {

	// num3Class := 2327
	// num4Class := 3615
	// num3 := 3
	// num4 := 4
	// fmt.Printf("3. Klasse %d: %d / %d = %f\n", num3Class, num3Class, num3, float64(num3)/float64(num3Class)*100)
	// fmt.Printf("4. Klasse %d: %d / %d = %f\n", num4Class, num4Class, num4, float64(num4)/float64(num4Class)*100)

	origin := vector.NewVec3(0.0, 0.0, 0.0)

	img := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	fmt.Printf("Image has size of %dx%d", img.Rect.Dx(), img.Rect.Dy())
	for x := img.Rect.Max.X - 1; x >= img.Rect.Min.X; x-- {
		for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {

			dr := float64(x+img.Rect.Min.X) / float64(img.Rect.Dx())
			dg := float64(y+img.Rect.Min.Y) / float64(img.Rect.Dy())

			ray := ray.NewRay(origin, cameraDirection(dr, dg))
			c := rayToColor(ray)

			img.SetRGBA(x, y, c)
		}
	}

	storeAsPNG(img, "image.png")
}
