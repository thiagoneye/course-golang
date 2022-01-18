/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/w_JW1bWT08s
Cap. 07.02.
*/

package main

import "fmt"

func main() {
    for j:= 65; j <= 90; j++ {
        for i := 0; i < 3; i++ {
            fmt.Printf("%#U\t", j)
        }
        fmt.Print("\n")
    }
}
