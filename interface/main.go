package main

import "fmt"

// Bkash
// Nagad
// Upay

// Tasks : SendMoney , CashOut

type PaymentProcessor interface {
	SendMoney(float64) string
	CashOut(float64) string
}

type Bkash struct {
	// Bkash Config
	BkashAPIKey string
}

type Nagad struct {
	NagadAPIKey string
	APISecrects string
}

func (b Bkash) SendMoney(amount float64) string {
	//dwdwdwdw
	//dwdwdwdw

	return "Send money successful "
}

func (b Bkash) CashOut(amount float64) string {
	//dwdwdwdw
	//dwdwdwdw

	return "Cashout successful "
}

func (n Nagad) SendMoney(amount float64) string {
	//dwdwdwdw
	//dwdwdwdw

	return "Send money successful "
}

func (n Nagad) CashOut(amount float64) string {
	//dwdwdwdw
	//dwdwdwdw

	return "Cashout successful "
}

type APIServer struct {
	paymentProcessor PaymentProcessor
}

func main() {
	BkashaAiServer := APIServer{
		paymentProcessor: Bkash{
			BkashAPIKey: "aaasasasasasa",
		},
	}
	res := BkashaAiServer.paymentProcessor.CashOut(100)
	fmt.Println(res)
	NagadAPIServer := APIServer{
		paymentProcessor: Nagad{
			APISecrects: "aaaaa",
			NagadAPIKey: "aaaaa",
		},
	}
	NagadAPIServer.paymentProcessor.SendMoney()

}
