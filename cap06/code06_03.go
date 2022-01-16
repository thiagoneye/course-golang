/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/QDaiqhTq3TA
Cap. 06.04.The for statement
*/

/*
Go não possui o comando "while".
Podemos simular um comando "while" com o comando "for".
*/

package main

import "fmt"

func main() {
    x := 0

    for x < 10 {
        fmt.Println(x)
        x++
    }
}
