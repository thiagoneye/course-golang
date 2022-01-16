/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/LBlrrV0iRKg
Cap. 06.07. The if statement
*/

package main

import "fmt"

func main() {
    x := 10
    if (x < 100) {
        fmt.Println("This is true!")
    }
    if !(x > 100) {
        fmt.Println("This is not true!")
    }
}
