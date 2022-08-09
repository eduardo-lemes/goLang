package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo  insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Não pode depositar", c.saldo
	}
}

func (c *ContaCorrente) Transferencia(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia > 0 && valorTransferencia < c.saldo {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaEduardo := ContaCorrente{"Eduardo", 567, 123456, 125.5}
	contaLemes := ContaCorrente{"Lemes", 857, 213213, 140.5}

	//fmt.Println(contaEduardo)
	//fmt.Println(contaLemes)

	//fmt.Println(contaEduardo.Sacar(500))
	//fmt.Println(contaEduardo)
	//status, valor := contaEduardo.Depositar(120)
	//fmt.Println(status, valor)

	statusTransferir := contaEduardo.Transferencia(100, &contaLemes)

	fmt.Println(statusTransferir)
	fmt.Println(contaEduardo)
	fmt.Println(contaLemes)
}
