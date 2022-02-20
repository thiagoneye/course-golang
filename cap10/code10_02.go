/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/KBFprVi_haM
Cap. 10.02. Embedded Fields
*/

/*
Os campos de um mesmo struct devem ser únicos.
*/

package main

import (
    "fmt"
)

type pessoa struct{
    nome string
    idade int
}

type profissional struct{
    pessoa
    titulo string
    salario int
}

func main() {
    pessoa1 := pessoa{
        nome: "Alfredo",
        idade: 30,
    }

    pessoa2 := profissional{
        pessoa: pessoa{
            nome: "Maria",
            idade: 31,
        },
        titulo: "Programadora",
        salario: 10000,
    }

    pessoa3 := profissional{pessoa{"José", 40}, "Cientista de Dados", 20000}

    fmt.Println(pessoa1)
    fmt.Println(pessoa1.nome)
    fmt.Println(pessoa2)
    fmt.Println(pessoa2.nome)
    fmt.Println(pessoa3)
}
