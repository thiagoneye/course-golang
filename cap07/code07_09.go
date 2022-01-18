/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/XxQQOHHQCwY
Cap. 07.09.
*/

package main

import "fmt"

func main() {
    esportefavorito := "paraquedismo"

    switch esportefavorito {
        case "futebol": fmt.Println("Esse é o esporte favorito dos brasileiros.")
        case "musculação": fmt.Println("Você irá ficar muito musculoso.")
        case "xadrez": fmt.Println("O esporte da estratégia!")
        case "paraquedismo": fmt.Println("Você gosta de ADRENALINA!!!")
        default: fmt.Println("Não sei o que dizer sobre esse esporte.")
    }
}
