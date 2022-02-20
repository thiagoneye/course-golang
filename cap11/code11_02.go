/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/uDajGJbXP6A
Cap. 11.02.
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

    pessoas := map[string][]string{
        "Gessinger": pessoa1.sabores,
        "Leindecker": pessoa2.sabores,
    }

    fmt.Println(pessoas)
    fmt.Println(pessoas["Gessinger"])
}
