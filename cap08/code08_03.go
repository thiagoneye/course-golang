/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/l6O8HWaoy_w
Cap. 08.03. Slice: for range
*/

package main

import "fmt"

func main() {
    slice := []string{"banana", "maça", "jaca"}
    slice = append(slice, "melancia")

    for indice, valor := range slice {
        fmt.Println("No índice", indice, "temos o valor", valor)
    }

    for _, valor := range slice {
        fmt.Printf("Um dos valores desse slice é %v.\n", valor)
    }
}
