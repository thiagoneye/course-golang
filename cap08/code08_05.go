/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/MbvABKiAABA
Cap. 08.05. Slice: Appending to a slice
*/

/*
... é chamado de unfurl, que funciona como um desempacotamento.
*/

package main

import "fmt"

func main() {
    slice1 := []int{1, 2, 3, 4}
    slice2 := []int{9, 10, 11, 12}

    slice3 := append(slice1, 5, 6, 7, 8)
    fmt.Println(slice3)

    slice4 := append(slice3, slice2...)
    fmt.Println(slice4)
}
