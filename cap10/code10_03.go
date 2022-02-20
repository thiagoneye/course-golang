/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/Y4MKS3gJQ9Q
Cap. 10.03. Anonymous Structures
*/

/*
Structs anônimos são "descartáveis".
*/

package main

import "fmt"

func main() {
    pessoa := struct{
        nome string
        idade int
    }{
        nome: "Mario",
        idade: 50,
    }

    fmt.Println(pessoa)
}
