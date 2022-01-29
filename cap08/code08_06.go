/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/IMO5_ancK9w
Cap. 08.06. Slice: Make
*/

/*
A função make é utilizada para criar uma slice com um array subjacente com o
objetivo de otimização de performance.
make([]type, len, cap)
*/

package main

import "fmt"

func main() {
    slice := make([]int, 5, 7)
    slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4
    fmt.Println(slice)

    slice = append(slice, 5)
    fmt.Printf("%v: %T\n", slice, slice)

    slice = append(slice, 6, 7, 8)
    fmt.Println(slice, len(slice), cap(slice))
}
