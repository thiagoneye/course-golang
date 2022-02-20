/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/uNnLSuRqBuY
Cap. 11.04.
*/

package main

import "fmt"

func main() {
    notebook := struct{
        marca string
        especificacoes map[string]int
        SO []string
    }{
        marca: "Acer",
        especificacoes: map[string]int{
            "memoria": 8,
            "ssd": 512,
        },
        SO: []string{"Windows", "Linux"},
    }

    fmt.Println(notebook)
}
