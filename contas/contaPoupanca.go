package contas

import "github.com/eduardo-lemes/goLang/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo  insuficiente."
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Não pode depositar", c.saldo
	}
}

func (c *ContaPoupanca) MostraSaldo() float64 {
	return c.saldo

}
