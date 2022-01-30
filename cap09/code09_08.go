/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/rfQkR1bH3qw
Cap. 09.08.
*/

package main

import "fmt"

func main() {
    dict := map[string][]string {
        "Maria José": []string{"Pintura", "Crochê"},
        "José Maria": []string{"Pesca", "Jornal"},
        "Ana Maria": []string{"Música", "Dança"},
    }

    for key, value := range dict {
        fmt.Printf("%v: %v\n", key, value)
    }
}
