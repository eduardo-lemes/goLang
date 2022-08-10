package main

import (
	"fmt"

	"github.com/eduardo-lemes/goLang/contas"
)

func main() {
	contaEdu := contas.ContaPoupanca{}
	contaEdu.Depositar(100)
	contaEdu.Sacar(50)

	fmt.Println(contaEdu.MostraSaldo())
}
