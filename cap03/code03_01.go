// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/5K17jFDXvWw
// Cap. 03.01.

package main

import "fmt"

func main() {
    x, y, z := 42, "James Bond", true

    fmt.Print(x, " ", y, " ", z)

    fmt.Println("\n\nx:", x)
    fmt.Println("y:", y)
    fmt.Println("z:", z)

    fmt.Printf("\nx (%T) = %v", x, x)
    fmt.Printf("\ny (%T) = %v", y, y)
    fmt.Printf("\nz (%T) = %v", z, z)
}
