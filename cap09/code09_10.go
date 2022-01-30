/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/2cx79-nwNQU
Cap. 09.10.
*/

package main

import "fmt"

func main() {
    dict := map[string][]string {
        "Maria José": []string{"Pintura", "Crochê"},
        "José Maria": []string{"Pesca", "Jornal"},
        "Ana Maria": []string{"Música", "Dança"},
    }

    delete(dict, "José Maria")

    for key, value := range dict {
        fmt.Printf("%v: %v\n", key, value)
    }
}
