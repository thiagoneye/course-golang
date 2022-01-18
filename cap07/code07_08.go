/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/mm3JM2ZG0us
Cap. 07.08.
*/

package main

import "fmt"

func main() {
    today := 15

    switch {
        case today < 15: fmt.Println("Ainda é início do mês.")
        case today == 15: fmt.Println("Passou exatamente a metade do mês.")
        case today > 15: fmt.Println("Já se passou mais da metade do mês.")
    }
}
