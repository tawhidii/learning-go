// package main

// import "fmt"

// /*
// Payment Processor : Bkash , Nagad , Upay
// Actions : CashOut, Bill Payment
// */
// type Account struct {
// 	balance int
// }

// type Processor interface {
// 	CashOut(string, int) (bool, error)
// 	BillPayment(string, int) (bool, error)
// }

// type PaymentProcessor struct {
// 	paymentProcessor Processor
// }

// type Bkash struct {
// 	Account Account
// }

// func (b *Bkash) CashOut(phoneNumber string, amount int) (bool, error) {
// 	b.Account.balance -= amount
// 	fmt.Println("CashOut success for " + phoneNumber + "!!")
// 	fmt.Println("Current Balance is : ", b.Account.balance)
// 	return true, nil
// }

// func (b *Bkash) BillPayment(serviceName string, amount int) (bool, error) {
// 	b.Account.balance -= amount
// 	fmt.Println("Bill payment success for " + serviceName + "!!")
// 	fmt.Println("Current Balance is : ", b.Account.balance)
// 	return true, nil
// }

// func main() {
// 	paymentProcessorBkash := PaymentProcessor{
// 		paymentProcessor: &Bkash{
// 			Account: Account{
// 				balance: 5000,
// 			},
// 		},
// 	}
// 	paymentProcessorBkash.paymentProcessor.BillPayment("electricity", 300)
// 	paymentProcessorBkash.paymentProcessor.CashOut("01616716072", 500)

// }
