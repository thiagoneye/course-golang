/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/o3yoYGWqrDE
Cap. 08.07. Slice: Multidimensional
*/

package main

import "fmt"

func main() {
    matrix := [][]int{
                []int{1, 2, 3},
                []int{4, 5, 6},
                []int{7, 8, 9},
    }

    fmt.Println(matrix)
    fmt.Println(matrix[2])
    fmt.Println(matrix[1][1])
}
