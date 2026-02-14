package main

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	df := dataframe.LoadStructs([]struct {
		Name string
		Age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
	})
	fmt.Println(df)
}
