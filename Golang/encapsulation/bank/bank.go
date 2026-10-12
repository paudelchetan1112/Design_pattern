package bank


import "fmt"
type BankAccount struct{
	balance float64
}
func NewBankAccount(initialBalance float64)*BankAccount{
	return &BankAccount{balance: initialBalance}
}
func(a *BankAccount) Deposite(amount float64){
	if amount<0{
		fmt.Println("Invalid amount")
		return
	}
	a.balance+=amount
}
func(a *BankAccount) Withdraw(amount float64){
		if amount<0{
		fmt.Println("Invalid amount")
		return
	}
if (a.balance-amount)<0{
	fmt.Println("Insufficient balance")
	return
}
a.balance-=amount
}
func(a BankAccount)Getbalance()float64{
	return a.balance
}
