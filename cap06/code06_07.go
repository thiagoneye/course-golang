/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/ZfCgoVaxMGE
Cap. 06.08. if, else if and else
*/

package main

import "fmt"

func main() {
    x := 5
    for x <= 15 {
        if x > 10 {
            fmt.Println("x é maior do que 10.")
        } else if x == 10 {
            fmt.Println("x é igual a 10.")
        } else {
            fmt.Println("x é menor do que 10.")
        }

        x++
    }
}
