//Programas executáveis iniciam pela pacote main
package main

/*
Os códigos em go são organizados em pacotes
e para usá-los é necessário declarar um ou vários imports
*/
import (
"fmt"
)

//A porta de entrada de um programa go é a função main
func main() {
	fmt.Print("Primeiro ")
	fmt.Print("Programa!")
}