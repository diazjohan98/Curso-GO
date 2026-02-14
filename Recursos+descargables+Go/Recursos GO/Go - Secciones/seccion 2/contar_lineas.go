package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("datos.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineas := 0
	for scanner.Scan() {
		lineas++
	}

	fmt.Println("Líneas:", lineas)
}
