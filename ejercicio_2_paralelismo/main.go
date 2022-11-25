package main

import (  
    "os"; "image"; "image/jpeg"; "image/color" ; "time"; "log" ; "sync"
    "path/filepath"; "fmt"; "strings"  
)

func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
}



func main() {  
  var equiz float64 
  fmt.Println("Operador Blending (valor x): ")
  fmt.Scanf("%f\n", &equiz) // AQUÍ VA EL VALOR DE X PARA EL OPERADOR BLENDING
  
    imgPath := "tigger.jpg"
	f, err := os.Open(imgPath)
	check(err)
	defer f.Close()

  imgPath2 := "plaza.jpg"
	f2, err := os.Open(imgPath2)
	check(err)
	defer f2.Close()

	img, _, err := image.Decode(f)

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	wImg := image.NewRGBA(rect)

  img2, _, err := image.Decode(f2)

	//size2 := img2.Bounds().Size()
	//rect2 := image.Rect(0, 0, size.X, size.Y)
	//wImg2 := image.NewRGBA(rect)   

    wg := new(sync.WaitGroup)

    start := time.Now()
 
    for y := 0; y < size.Y; y++ {   
        wg.Add(1)       
        y := y
        // estamos creado un go routine por cada fila 
        // NO es lo recomendable, porque vamos a tener mucho gorutines
        // lo recomendable seria tener solo unos cuantos go routines (max 5)
        go func (){
            for x := 0; x < size.X; x++ {      
                pixel := img.At(x, y)
				originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

        pixel2 := img2.At(x, y)
				originalColor2 := color.RGBAModel.Convert(pixel2).(color.RGBA)

				// Compensa los colores un poco, ajústalo a tu gusto
				r := float64(originalColor.R)// * 0.92126 
				g := float64(originalColor.G)// * 0.97152 
				b := float64(originalColor.B)// * 0.90722
        //a := float64(originalColor.A)

        r2 := float64(originalColor2.R)// * 0.92126 
				g2 := float64(originalColor2.G)// * 0.97152 
				b2 := float64(originalColor2.B)// * 0.90722
        //a2 := float64(originalColor.A)
				// suma
        //total := uint8(r+r2+g+g2+b+b2)
				sumR := uint8(equiz*r+(1-equiz)*r2)
        sumG := uint8(equiz*g+(1-equiz)*g2)
        sumB := uint8(equiz*b+(1-equiz)*b2)
                c := color.RGBA{ 
                    R: sumR, G: sumG, B: sumB, A: originalColor.A, 
                } 
                wImg.Set(x, y, c)                                          
            } 
            defer wg.Done()   
        }()
    }
    wg.Wait()
            
    elapsed := time.Since(start)
    log.Printf("El operador blending tomó %s", elapsed)

    ext := filepath.Ext(imgPath)  
    name := strings.TrimSuffix(filepath.Base(imgPath), ext)  
    newImagePath := fmt.Sprintf("%s/%s_blending%s", filepath.Dir(imgPath), name, ext)  
    fg, err := os.Create(newImagePath)  
    defer fg.Close()
    check(err)
    err = jpeg.Encode(fg, wImg, nil)
    check(err)
}