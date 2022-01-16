/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/EL9wo6Zrz9o
Cap. 06.03. Nested loop
*/

package main

import "fmt"

func main() {
    for hour := 0; hour < 12; hour ++ {
        fmt.Println("Hora:", hour)

        for minute := 0; minute < 60; minute ++ {
            fmt.Print(" ", minute)
        }

        fmt.Print("\n")
    }
}
