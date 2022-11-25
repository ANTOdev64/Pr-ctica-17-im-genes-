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
  "bufio"
  "sync"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func writeTxt(lines []float64, path string) error { 
 file, err := os.Create(path)
 if err != nil {
   return err 
 } 
 defer file.Close() 
  
  w := bufio.NewWriter(file) 
 for _, line := range lines {
   fmt.Fprintln(w, line) 
 } 
 return w.Flush() 
}

func main() {
  var R []float64
  var G []float64
  var B []float64
  
  imgPath := "arequipa.jpg"
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

	img, _, err := image.Decode(f)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

  wg := new(sync.WaitGroup)

	start := time.Now()

	
		for y := 0; y < size.Y; y++ {
      wg.Add(1)       
        y := y
			// y ahora recorre todo esto x y
			go func(){
        for x := 0; x < size.X; x++ {

				pixel := img.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

				// Compensa los colores un poco, ajústalo a tu gusto
				r := float64(originalColor.R)// * 0.92126 
				g := float64(originalColor.G)// * 0.97152 
				b := float64(originalColor.B)// * 0.90722
        
        R = append(R, r)
        G = append(G, g)
        B = append(B, b)
        
				c := color.RGBA{ R: originalColor.R, G: originalColor.G, B: originalColor.B, A: originalColor.A, }
				wImg.Set(x, y, c)
			}
      defer wg.Done() 
    }()
	}
  wg.Wait()
  
	elapsed := time.Since(start)
	log.Printf("El proseso tomó %s", elapsed)

	ext := filepath.Ext(imgPath)
	name := strings.TrimSuffix(filepath.Base(imgPath), ext)
	newImagePath := fmt.Sprintf("%s/%s_prosesado%s", filepath.Dir(imgPath), name, ext)
	fg, err := os.Create(newImagePath)
	defer fg.Close()
	check(err)
	err = jpeg.Encode(fg, wImg, nil)
	check(err)

  writeTxt(R, "Rojo.txt")
  writeTxt(G, "Verde.txt")
  writeTxt(B, "Azul.txt")
}