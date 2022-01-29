/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/MMzTlWZ9Gjw
Cap. 08.02. Slice: Compound Literal
*/

/*
Os slices são mais flexíveis que do que os arrays.
Os slices são tipos de dados compostos.

Os arrays possuem um número finito de elementos pré-fixado.
Os slices possuem um número variável de elementos.
*/

package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println(array)

    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    slice2 := append(slice, 6)
    fmt.Println(slice2)

    slice[3] = 333
    fmt.Println(slice)
}
