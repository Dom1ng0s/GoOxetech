package main

import "fmt"

type ContaBancaria struct {
	titular string
	Saldo   float64
}

func (c ContaBancaria) ExibirSaldo() float64 {
	return c.Saldo
}

func (c *ContaBancaria) Depositar(deposito float64) {
	c.Saldo += deposito
	fmt.Printf("Depósito de %.2f realizado. Novo Saldo: %.2f\n", deposito, c.ExibirSaldo())
}

func (c *ContaBancaria) Sacar(saque float64) bool {
	if saque > c.Saldo {
		fmt.Printf("Saque de %.2f não permitido. Saldo insuficiente (%.2f).\n", saque, c.Saldo)
		return false
	} else {
		c.Saldo -= saque
		fmt.Printf("Saque de %.2f realizado. Novo Saldo: %.2f\n", saque, c.ExibirSaldo())
		return true
	}
}

func main() {
	minhaConta := ContaBancaria{titular: "João", Saldo: 1500.50}

	fmt.Println("--- INÍCIO DOS TESTES ---")
	fmt.Printf("Saldo Inicial da conta de %s: %.2f\n\n", minhaConta.titular, minhaConta.ExibirSaldo())

	minhaConta.Depositar(250.00)

	minhaConta.Sacar(100.00)

	minhaConta.Sacar(2000.00)

	saldoFinal := minhaConta.ExibirSaldo()
	fmt.Printf("\nSaldo Final da conta: %.2f\n", saldoFinal)
}
