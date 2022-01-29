/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/clobcQqAgos
Cap. 08.09. Maps: Introduction
*/

/*
Map é uma estrutura não ordenada do tipo chave-valor.
*/

package main

import "fmt"

func main() {
    amigos := map[string]int {
        "Alfredo": 53638940,
        "Joana": 93827491,
    }

    fmt.Println(amigos)
    fmt.Println(amigos["Joana"])

    // Adicionando elemento
    amigos["Gopher"] = 46380487
    fmt.Println(amigos)

    // Verificando se elemento existe
    fmt.Println(amigos["Romário"])

    pessoa, existe := amigos["Romário"]
    fmt.Println(pessoa, existe)
}
