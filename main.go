package main

import (
	"fmt"

	"./contas"
)

func main() {
	contaEduardo := contas.ContaCorrente{Titular: "Eduardo", NumeroAgencia: 567, NumeroConta: 123456, Saldo: 125.5}
	contaLemes := contas.ContaCorrente{Titular: "Lemes", NumeroAgencia: 857, NumeroConta: 213213, Saldo: 140.5}

	//fmt.Println(contaEduardo)
	//fmt.Println(contaLemes)

	//fmt.Println(contaEduardo.Sacar(500))
	//fmt.Println(contaEduardo)
	//status, valor := contaEduardo.Depositar(120)
	//fmt.Println(status, valor)

	statusTransferir := contaEduardo.Transferencia(-300, &contaLemes)

	fmt.Println(statusTransferir)
	fmt.Println(contaEduardo)
	fmt.Println(contaLemes)
}
