/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/v6HnDiPyynA
Cap. 06.10. The switch statement (part 2)
*/

package main

import "fmt"

var x interface{}

func main() {
    x = true

    switch x.(type) {
        case int: fmt.Println("It's a int")
        case float64: fmt.Println("It's a float")
        case bool: fmt.Println("It's a bool")
        case string: fmt.Println("It's a string")
    }
}
