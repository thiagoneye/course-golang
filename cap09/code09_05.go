/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/gA1cTnWjPaU
Cap. 09.05.
*/

package main

import "fmt"

func main() {
    x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

    y := make([]int, len(x))
    copy(y, x)

    x = append(x[:3], x[7:]...)
    
    fmt.Println("x:", x)
    fmt.Println("y:", y)
}
