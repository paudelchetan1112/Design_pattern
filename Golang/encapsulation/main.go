package main

import (
	"fmt"
	"bankapp/bank"
)
func main(){
	
	myAccount:=bank.NewBankAccount(1000)
	fmt.Println(myAccount)
	fmt.Println(myAccount.Getbalance())
	myAccount.Withdraw(200)

	myAccount.Deposite(100)


fmt.Println(myAccount.Getbalance())

}
