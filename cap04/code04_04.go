/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/AcyS0_BAy7U
Cap. 04.05. Strings
*/

/*
Comments:
Strings are immutable
*/

package main

import "fmt"

func main() {
    s := "Hello!"
    fmt.Printf("%v: %T\n", s, s)

    fmt.Println("")

    for _, v := range s {
        fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
    }

    fmt.Println("")

    sb := []byte(s)
    fmt.Printf("%v: %T\n", sb, sb)

    fmt.Println("")

    for _, v := range sb {
        fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
    }

    fmt.Println("")

    s = `Hello!

            This is a literal string.`
    fmt.Printf("%v\n", s)

    fmt.Println("")

    sb = []byte(s)
    fmt.Printf("%v: %T\n", sb, sb)
}
