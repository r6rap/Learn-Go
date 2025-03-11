package latihansoal

import(
	"fmt"
)

type BankAccount struct{
	AccountNumber, HolderName string
	Balance float64
}

	  //menggunakan pointer agar dinamis
func (b *BankAccount) deposit(amount float64) error {
	if amount <= 10000{
		return fmt.Errorf("Deposit minimal 10.000")
	}

	fmt.Printf("Saldo awal: %.2f\n", b.Balance)
	b.Balance += amount
	fmt.Printf("Total deposit: %.2f\n", amount)
	fmt.Printf("Saldo sekarang: %.2f\n", b.Balance)

	return nil
}

func (b *BankAccount) withdraw(amount float64) error {
	if amount < 0{
		err := fmt.Errorf("Jumlah withdraw tidak valid")
		fmt.Println(err)
		return err
	}

	if b.Balance < amount{
		err := fmt.Errorf("Saldo tidak cukup untuk melakukan withdraw")
		fmt.Println("Error", err)
		return err
	}
	fmt.Printf("Saldo awal: %.2f\n", b.Balance)
	b.Balance -= amount
	fmt.Printf("Total withdraw: %.2f\n", amount)
	fmt.Printf("Saldo sekarang: %.2f\n", b.Balance)

	return nil
}

func (k *BankAccount) transfer(target *BankAccount, amount float64) error {
	if amount < 10000{
		return fmt.Errorf("Jumlah minimal transfer adalah 10.000")
	}

	k.withdraw(amount) //ambil / wd uang dari balance kita
	dp := target.deposit(amount) // balance dari target ditambah dengan amount
	if dp == nil{
		fmt.Println("Berhasil transfer ke", target.AccountNumber, "sebesar", amount)
		fmt.Printf("Saldo target %.2f\n", target.Balance)
		fmt.Printf("Saldo sekarang %.2f\n", k.Balance)
	}

	return nil
}

func Testing() {
	B1 := &BankAccount{AccountNumber: "6784653", HolderName: "Rap", Balance: 1000000}
	B2 := &BankAccount{AccountNumber: "6758943", HolderName: "Pip", Balance: 500000}

	err := B1.transfer(B2, 50000)
	if err != nil{
		fmt.Println(err)
	}
	// err := B1.withdraw(200000)
	// if err != nil{
	// 	fmt.Println(err)
	// }

	// err = B1.deposit(100000)
	// if err != nil{
	// 	fmt.Println(err)
	// }
}