/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/haYnFVssUr4
Cap. 07.07.
*/

package main

import "fmt"

func main() {
    today := 18
    if (today < 15) {
        fmt.Println("Ainda é início do mês.")
    } else if (today == 15) {
        fmt.Println("Passou exatamente a metade do mês.")
    } else {
        fmt.Println("Já se passou mais da metade do mês.")
    }
}
