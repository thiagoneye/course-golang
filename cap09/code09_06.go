/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/KOhQCm8f8AE
Cap. 09.06.
*/

package main

import "fmt"

func main() {
    values := []string{"AC", "AL", "AP", "AM", "BA", "CE", "ES", "GO", "MA",
              "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS",
              "RO", "RR", "SC", "SP", "SE", "TO"}

    states := make([]string, len(values))

    for i, v := range values {
        states[i] = v
    }

    fmt.Println(states)
    fmt.Printf("T: %T, L: %v and C: %v", states, len(states), cap(states))
}
