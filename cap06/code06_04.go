/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/eGdB8FMCZ0s
Cap. 06.05. break and continue
*/

package main

import "fmt"

func main() {
    for x := 0; x < 10; x++ {
        if (x < 5) {
            fmt.Println(x)
            continue
        } else {
            break
        }
        fmt.Println("Essa função não será executada por causa do continue.")
    }
}
