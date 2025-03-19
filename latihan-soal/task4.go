package latihansoal

import(
	"fmt"
)

type BankAccount struct{
	AccountNumber, HolderName string
	Balance float64
}

	//menggunakan method pointer agar dinamis
func (b *BankAccount) deposit(amount float64) error {
	if amount <= 10000{
		return fmt.Errorf("deposit minimal 10.000")
	}

	fmt.Printf("Saldo awal: %.2f\n", b.Balance)
	b.Balance += amount
	fmt.Printf("Total deposit: %.2f\n", amount)
	fmt.Printf("Saldo sekarang: %.2f\n", b.Balance)

	return nil
}

func (b *BankAccount) withdraw(amount float64) error {
	if amount < 0{
		err := fmt.Errorf("jumlah withdraw tidak valid")
		fmt.Println(err)
		return err
	}

	if b.Balance < amount{
		err := fmt.Errorf("saldo tidak cukup untuk melakukan withdraw")
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
		return fmt.Errorf("jumlah minimal transfer adalah 10.000")
	}

	if err := k.withdraw(amount); err != nil {
		return err
	} //ambil / wd uang dari balance kita sebanyak amount
	
	dp := target.deposit(amount) // balance dari target ditambah dengan amount yang sudah kita wd dari balance kita
	if dp == nil{
		fmt.Println("Berhasil transfer ke", target.AccountNumber, "sebesar", amount)
		fmt.Printf("Saldo target %.2f\n", target.Balance)
		fmt.Printf("Saldo sekarang %.2f\n", k.Balance)
	}

	return nil
}

	//menggunakan method value karena hanya menghitung bunga berdasarkan saldo dari instance
func (b BankAccount) sukuBunga() (float64, error) {
	if b.Balance <= 10000 {
		return 0, fmt.Errorf("saldo harus di atas 10.000")
	}

	bunga := b.Balance * 0.04

	return bunga, nil
}

	//menggunakan method value karena hanya menampilkan info akun berdasarkan instance
func (b BankAccount) display() (string, string, float64) {
	return b.AccountNumber, b.HolderName, b.Balance //mengembalikan sesuai dengan instancenya
}

func Testing() {
	B1 := BankAccount{AccountNumber: "6784653", HolderName: "Rap", Balance: 1000000}
	B2 := BankAccount{AccountNumber: "6758943", HolderName: "Pip", Balance: 500000}


	err := B1.transfer(&B2, 50000)
	if err != nil{
		fmt.Println(err)
	}

	hsl, err := B2.sukuBunga()
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Bunga pertahun adalah", hsl)
	}

	accNum, name, balance := B1.display()
	fmt.Println("No rekening:", accNum)
	fmt.Println("Nama rekening:", name)
	fmt.Println("Saldo rekening:", balance)
}
