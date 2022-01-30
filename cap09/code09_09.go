/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/AKPdowl7tsw
Cap. 09.09.
*/

package main

import "fmt"

func main() {
    dict := map[string][]string {
        "Maria José": []string{"Pintura", "Crochê"},
        "José Maria": []string{"Pesca", "Jornal"},
        "Ana Maria": []string{"Música", "Dança"},
    }

    dict["João Paulo"] = []string{"Grafite", "Pedalada"}

    for key, value := range dict {
        fmt.Printf("%v: %v\n", key, value)
    }
}
