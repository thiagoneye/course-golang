// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/DGHYibXY6Qk
// Cap. 02.08.
// Criando seu próprio tipo

package main

import "fmt"

type hotdog int
var a hotdog

func main() {
    fmt.Printf("%T", a)
}
