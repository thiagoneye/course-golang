// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/QT7YvrEWMGE
// https://github.com/ellenkorbes/aprendago
// Cap. 02.04.

package main

import (
	"fmt"
)

func main (){
		x:= 10
		y := "Olá, bom dia."

		fmt.Printf("x: %v, %T\n", x, x)
		fmt.Printf("y: %v, %T\n\n", y, y)

		x, z := 20, 30
		fmt.Printf("x: %v, %T\n", x, x)
		fmt.Printf("z: %v, %T\n", z, z)
}
