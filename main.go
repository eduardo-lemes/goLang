package main

import (
	"fmt"

	"github.com/eduardo-lemes/goLang/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaEdu := contas.ContaPoupanca{}
	contaEdu.Depositar(100)
	PagarBoleto(&contaEdu, 50)

	fmt.Println(contaEdu.MostraSaldo())

	contaLemes := contas.ContaCorrente{}
	contaLemes.Depositar(300)
	PagarBoleto(&contaLemes, 100)

	fmt.Println(contaLemes.MostraSaldo())
}
