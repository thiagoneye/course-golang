/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/7a6I-GnqtSE
Cap. 08.10. Maps: range and deleting
*/

package main

import "fmt"

func main() {
    qualquercoisa := map[int]string{
        123: "muito legal",
        98: "menos legal um pouquinho",
        983: "esse é massa",
        18: "idade de ir para as festas",
    }

    fmt.Println(qualquercoisa)

    for key, value := range qualquercoisa {
        fmt.Println(key, value)
    }

    delete(qualquercoisa, 983)
    fmt.Println(qualquercoisa)
}
