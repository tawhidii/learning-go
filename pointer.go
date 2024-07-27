// package main

// import "fmt"

// type Account struct {
// 	balance int
// }

// // func withdrawCash(account *Account, amount int) {
// // 	account.balance -= amount
// // 	fmt.Println("Withdraw amount is : ", amount, "Current Balance is: ", account.balance)
// // }

// // using function reciever
// func (account *Account) withDrawCash(amount int) {
// 	account.balance -= amount
// 	fmt.Println("Withdraw amount is : ", amount, "Current Balance is: ", account.balance)
// }

// func main() {
// 	account := &Account{
// 		balance: 10000,
// 	}

// 	account.withDrawCash(100)

// 	fmt.Println("Total Balance: ", account.balance)

// }
