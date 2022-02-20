/*
https://github.com/ellenkorbes/aprendago
https://www.youtube.com/c/AprendaGo/playlists
https://youtu.be/ji14zPQgmN8
Cap. 11.03.
*/

package main

import "fmt"

type veiculo struct{
    portas int
    cor string
}

type caminhonete struct{
    veiculo
    tracaoQuatroRodas bool
}

type sedan struct{
    veiculo
    modeloLuxo bool
}

func main() {
    carro1 := caminhonete{
        veiculo: veiculo{portas: 6, cor: "preto"},
        tracaoQuatroRodas: true,
    }

    carro2 := sedan{
        veiculo: veiculo{portas: 4, cor: "vinho"},
        modeloLuxo: false,
    }

    fmt.Println(carro1)
    fmt.Println(carro1.cor)
    fmt.Println(carro2)
    fmt.Println(carro2.cor)
}
