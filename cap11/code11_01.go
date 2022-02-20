/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/OQMOIcZ-ShY
Cap. 11.01.
*/

package main

import "fmt"

type pessoa struct{
    nome string
    sobrenome string
    sabores []string
}

func main() {
    pessoa1 := pessoa{
        nome: "Humberto",
        sobrenome: "Gessinger",
        sabores: []string{"Flocos", "Creme com Passas"},
    }

    pessoa2 := pessoa{
        nome: "Duca",
        sobrenome: "Leindecker",
        sabores: []string{"Oreo", "Morango", "Chocolate Africano"},
    }

    fmt.Println(pessoa1)
    fmt.Println(pessoa2)
}
