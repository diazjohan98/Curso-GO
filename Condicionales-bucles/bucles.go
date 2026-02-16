package main

import "fmt"

func main() {

	// * bucle clasico
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración:", i)
	}

	// * bucle while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	numeros := []int{1, 2, 3, 4, 5}

	for i, num := range numeros {
		fmt.Println("Indice:", i, "Valor:", num)
	}

	edades := map[string]int{
		"Johan": 27,
		"Maria": 30,
	}

	for nombre, edad := range edades {
		fmt.Println(nombre, "tiene", edad, "años")
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
