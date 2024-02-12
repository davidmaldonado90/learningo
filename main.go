package main

import (
	"fmt"

	"github.com/dmaldonado/learningo/ejercicios"
)

func main() {
	numero, texto := ejercicios.ConvNum("f")
	fmt.Println(numero)
	fmt.Println(texto)
}
