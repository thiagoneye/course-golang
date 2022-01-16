/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/WFFidtqPfhk
Cap. 06.09. The switch statement
*/

package main

import "fmt"

func main() {
    for i := 0; i <= 10; i++ {
        switch {
            case i < 5:
                fmt.Println("i é menor do que 5.")
            case i == 5:
                fmt.Println("i é igual a 5.")
            case i > 5:
                fmt.Println("i é maior do que 5.")
        }
    }

    expr := ""

    switch expr {
        case "exp":
            fmt.Println("\nexp")
        case "expres":
            fmt.Println("\nexpres")
        case "expression":
            fmt.Println("\nexpression")
        default:
            fmt.Println("\nNone of the cases.")
    }

    x := 3

    switch x {
        case 1, 2:
            fmt.Println("1 ou 2")
        case 3, 4:
            fmt.Println("3 ou 4")
        case 5, 6:
            fmt.Println("5 ou 6")
    }

    switch {
        case (x == 1), (x == 2):
            fmt.Println("1 ou 2")
        case (x == 3), (x == 4):
            fmt.Println("3 ou 4")
    }
}
