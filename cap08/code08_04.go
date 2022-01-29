/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/G0rxcnojV_U
Cap. 08.04. Slice: Slicing or deleting from a slice
*/

package main

import "fmt"

func main() {
    sabores := []string{"pepperoni", "mozzarella", "abacaxi", "4 queijos", "marguerita"}
    fatias := sabores[1:3]
    fmt.Println(fatias)

    sabores = append(sabores[:2], sabores[3:]...)
    fmt.Println(sabores)
}
