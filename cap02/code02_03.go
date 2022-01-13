// https://github.com/ellenkorbes/aprendago
// https://www.youtube.com/c/AprendaGo/videos
// https://youtu.be/IjDiONSi-tI
// Cap. 02.07.
// O Pacote "fmt"

package main

import "fmt"

func main() {
    x := "Oi!"
    y := "Oi, bom dia!\nComo vai você?\tEspero que tudo bem." // Interpreted string literals
    z := `Oi, bom dia!\nComo vai você?\tEspero que tudo bem.` // Raw string literals

    fmt.Print(x)   // SEM a quebra de linha no final
    fmt.Println(y) // COM a quebra de linha no final
    fmt.Println(z)

    x = "Oi!"
    y = "Meu nome é Thiago."

    xy := fmt.Sprint(x,y) // Converte as entradas em strings SEM adição de espaço
    fmt.Println(xy)

    xy = fmt.Sprintln(x,y) // Converte as entradas em strings COM adição de espaço
    fmt.Println(xy)

    fmt.Printf("%v: %T", x, x) // Formata de acordo com um especificador de formato e grava na saída padrão

    const name, age = "Kim", 22
	  s := fmt.Sprintf("\n%s is %d years old.\n", name, age) // Formata de acordo com um especificador de formato e retorna a string resultante

    fmt.Print(s)
}
