package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo  insuficiente."
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	podeDepositar := valorDeposito > 0
	if podeDepositar {
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.Saldo
	} else {
		return "Não pode depositar", c.Saldo
	}
}

func (c *ContaCorrente) Transferencia(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia > 0 && valorTransferencia < c.Saldo {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}
}
