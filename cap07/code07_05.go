/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/pIVF9vf2wAc
Cap. 07.05.
*/

package main

import "fmt"

func main() {
    var j int

    for i := 10; i <= 100; i++ {
        j = (i % 4)
        fmt.Printf("%v %% 4 = %v\n", i, j)
    }
}
