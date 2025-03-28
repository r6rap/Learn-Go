/*	Create a simple e-wallet system */

package latihansoal

import (
	"fmt"
)

type wallet struct{
	Number int64
	HolderName string
	Balance float64
	Transactions []string
}
type payment struct{
	company string
	pay float64
}

// function ga perfect
func (w *wallet) register(no int64, name string) (*wallet, error) {
	if w.Number != 0{
		return w, fmt.Errorf("no sudah terdaftar")
	}

	newWallet := wallet{
		Number: no,
		HolderName: name,
		Balance: 0,
	}

	return &newWallet, nil
}

func (t *wallet) topup(amount float64) error {
	if amount < 10000{
		return fmt.Errorf("jumlah minimal topup Rp. 10.000")
	}

	t.Balance += amount

	t.Transactions = append(t.Transactions, fmt.Sprintf("Topup sebesar: +%.2f", amount))

	return nil
}

func (p *wallet) payment(company string, amount float64) error {
	if p.Balance < amount {
		return fmt.Errorf("tidak bisa membayar, saldo tidak cukup")
	}
	p.Balance -= amount

	p.Transactions =  append(p.Transactions, fmt.Sprintf("Pembayaran %s sebesar: -%.2f", company, amount))

	return nil
}

func (tf *wallet) transfer(target *wallet, amount float64) error {
	if tf.Number == target.Number {
		return fmt.Errorf("tidak bisa transfer ke nomor sendiri")
	}
	if amount > tf.Balance {
		return fmt.Errorf("saldo tidak cukup untuk melakukan transfer")
	}

	tf.Balance -= amount
	tf.Transactions = append(tf.Transactions, fmt.Sprintf("Transfer ke %s -%.2f", target.HolderName, amount))

	target.Balance += amount
	target.Transactions = append(target.Transactions, fmt.Sprintf("Transfer dari %s +%.2f", tf.HolderName, amount))

	return nil
}

func (b *wallet) monthlyInterest() (float64, error) {
	if b.Balance < 1000000 {
		return 0, fmt.Errorf("saldo anda di bawah Rp 1.000.000, tidak bisa mendaoatkan bunga perbulan")
	}

	monthlyInterest := b.Balance * 0.02

	b.Balance += monthlyInterest

	b.Transactions = append(b.Transactions, fmt.Sprintf("Bunga bulanan Rp %.2f", monthlyInterest))

	return monthlyInterest, nil
}

func (info wallet) information() {
	fmt.Println("Nomor akun", info.Number)
	fmt.Println("Username akun", info.HolderName)
	fmt.Printf("Saldo: %.2f\n", info.Balance)
}

func (h *wallet) historyTransaction() error {
	if len(h.Transactions) == 0 {
		return fmt.Errorf("tidak ada transaksi")
	}

	fmt.Println("Riwayat Transaksi:")
	for i, v := range h.Transactions {
		fmt.Println(i+1, v)
	}

	fmt.Printf("Saldo sekarang: Rp %.2f", h.Balance)

	return nil
}

func P_Wallet() {
	netflix := payment{company: "netflix", pay: 69000}
	// spotify := payment{company: "spotify", pay: 55000}

	Acc := wallet{}
	
	Data1, Err := Acc.register(6289548392784, "Rafif Rizal")
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Println("Registrasi berhasil")
		fmt.Println("Nomor:",Data1.Number)
		fmt.Println("Nama:",Data1.HolderName)
		fmt.Printf("Saldo: %.2f\n", Data1.Balance)
	}

	Data2, Err := Acc.register(6287473846384, "Edbert Julian")
	if Err != nil {
		fmt.Println(Err)
	} else {
		fmt.Println("Registrasi berhasil")
		fmt.Println("Nomor:",Data2.Number)
		fmt.Println("Nama:",Data2.HolderName)
		fmt.Printf("Saldo: %.2f\n", Data2.Balance)
	}

	if err := Data1.topup(50000); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Topup berhasil, saldo dari akun %d sekarang adalah: %.2f\n",Data1.Number, Data1.Balance)
	}

	var amount float64 = 30000
	if err := Data1.transfer(Data2, amount); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Berhasil tf sebesar %.2f ke nomor %d\n", amount, Data2.Number)
		fmt.Printf("Saldo sekarang: %.2f\n", Data1.Balance)
		fmt.Printf("Saldo %s sekarang: %.2f\n", Data2.HolderName, Data2.Balance)
	}

	Data1.topup(100000)

	if err := Data1.payment(netflix.company, netflix.pay); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil membayar")
		fmt.Printf("Saldo sekarang: %.2f\n", Data1.Balance)
	}

	bunga, err := Data1.monthlyInterest()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Bunga anda setiap bulan: %.2f\n", bunga)
	}

	if err := Data1.historyTransaction(); err != nil {
		fmt.Println(err)
	}
}