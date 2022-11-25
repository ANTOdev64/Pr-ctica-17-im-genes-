package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
  
  imgPath := "leon.jpg"
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

  imgPath2 := "arequipa.jpg"
	f2, err := os.Open(imgPath2)
	check(err)
	defer f2.Close()

	img, _, err := image.Decode(f)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

  img2, _, err := image.Decode(f2)

	start := time.Now()

	func() {
		for x := 0; x < size.X; x++ {
			// y ahora recorre todo esto x y
			for y := 0; y < size.Y; y++ {

				pixel := img.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

        pixel2 := img2.At(x, y)
				originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)
				// Compensa los colores un poco, ajústalo a tu gusto
				r := float64(originalColor.R)// * 0.92126 
				g := float64(originalColor.G)// * 0.97152 
				b := float64(originalColor.B)// * 0.90722
        a := float64(originalColor.A)

        r2 := float64(originalColor2.R)// * 0.92126 
				g2 := float64(originalColor2.G)// * 0.97152 
				b2 := float64(originalColor2.B)// * 0.90722
        a2 := float64(originalColor2.A)
        
				sumR := uint8((r+r2)/2)
        sumG := uint8((g+g2)/2)
        sumB := uint8((b+b2)/2)
        sumA := uint8((a+a2)/2)
        
				c := color.RGBA{ R: sumR, G: sumG, B: sumB, A: sumA, }
				wImg.Set(x, y, c)
			}
		}
	}()
  
	elapsed := time.Since(start)
	log.Printf("La adición tomó %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_adicion%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)
}