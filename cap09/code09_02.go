/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/xpeExQ5C5S8
Cap. 09.02.
*/

package main

import "fmt"

func main() {
    slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

    for i, v := range slice {
        fmt.Printf("No índice %v, temos o valor %v.\n", i, v)
    }

    fmt.Printf("%v: %T", slice, slice)
}
