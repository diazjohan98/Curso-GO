package main

import "fmt"

type Persona struct {
    Nombre string
    Edad   int
}

func main() {
    var edad int = 30
    var salario float64 = 3500.50
    nombre := "Jorge"
    var activo bool = true

    fmt.Println(edad, salario, nombre, activo)

    numeros := []int{1, 2, 3, 4, 5}
    fmt.Println(numeros)
    numeros = append(numeros, 6)
    fmt.Println(numeros) 

    for i, num := range numeros {
        fmt.Println("Posición:", i, "Valor:", num)
    }

    p := Persona{Nombre: "Jorge", Edad: 30}
    fmt.Println(p)
    fmt.Println(p.Nombre)
    fmt.Println(p.Edad)

    edades := map[string]int{
        "Jorge": 30,
        "Ana":   25,
    }

    fmt.Println(edades)
    fmt.Println("Edad de Ana:", edades["Ana"])
    edades["Luis"] = 40
    fmt.Println(edades)


}
