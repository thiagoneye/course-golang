/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/EaOGcmXo4F8
Cap. 10.01. Struct
*/

package main

import "fmt"

type pessoa struct {
    nome string
    sobrenome string
    fumante bool
}

func main() {
    pessoa1 := pessoa{
        nome: "John",
        sobrenome: "Lennon",
        fumante: true,
    }
    pessoa2 := pessoa{"Juliana", "Paes", false}

    fmt.Println(pessoa1)
    fmt.Println(pessoa2)
}
