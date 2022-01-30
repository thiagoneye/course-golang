/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/f4lvrtXsGIM
Cap. 09.07.
*/

package main

import "fmt"

func main() {
    slice := [][]string{
      []string{"Maria", "José", "Pintura"},
      []string{"José", "Maria", "Pesca"},
      []string{"João", "Paulo", "Música"},
    }

    for _, v := range slice {
        fmt.Println(v)
    }
}
