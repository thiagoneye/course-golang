/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/hu0qcmbhH8s
Cap. 06.06. Using ASCII
*/

package main

import "fmt"

func main() {
    for i := 33; i <= 122; i++ {
        fmt.Printf("%d \t%b \t%#x \t%#U\n", i, i, i, i)
    }

    for i := 33; i <= 122; i++ {
        fmt.Printf("%v\n", string(i))
    }
}
