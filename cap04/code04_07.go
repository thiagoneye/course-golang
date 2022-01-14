/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/1IduyaGMO3g
Cap. 04.08. Iotas
*/

package main

import "fmt"

const (
    a = iota
    _ = iota
    c = iota
)

const (
    x = iota
    y = iota
    z = iota
)

const (
    e = 10
    f = iota
    g = iota
)

const (
    i = iota
    j
    k
)

const (
    l = iota*10
    m
    n
)

const (
    o = iota + 5
    p
    q
)

func main() {
    fmt.Println(a, c)
    fmt.Println(x, y, z)
    fmt.Println(e, f, g)
    fmt.Println(i, j, k)
    fmt.Println(l, m, n)
    fmt.Println(o, p, q)
}
