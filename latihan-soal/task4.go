package latihansoal

import(
	"fmt"
)

type BankAccount struct{
	AccountName, HolderName string
	Balance float64
}

func (b BankAccount) deposit(amount float64) {
	total := b.Balance + amount
	fmt.Println("Saldo awal:", b.Balance)
	fmt.Println("Deposit saldo sebesar:", amount)
	fmt.Println("Total saldo", total)
}

func (b BankAccount) withdraw(amount float64) error {
	if b.Balance < amount{
		return fmt.Errorf("Saldo tidak cukup untuk melakukan withdraw")
	}

	wd := b.Balance - amount
	fmt.Println("Saldo awal", b.Balance)
	fmt.Println("Total withdraw", amount)
	fmt.Println("Saldo sekarang", wd)

	return nil
}

func Testing() {
	B1 := BankAccount{AccountName: "Rap", HolderName: "BCA", Balance: 1000000}
	err := B1.withdraw(200000)
	if err != nil{
		fmt.Println(err)
	}

	//B1.deposit(100000)
}