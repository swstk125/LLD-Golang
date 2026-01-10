package main

import "fmt"

type Swiggy struct {
	walletManager      *WalletManager
	transactionManager *TransactionManager
	refundManager      *RefundManager
}

func NewSwiggyService() *Swiggy {
	return &Swiggy{
		walletManager:      NewWalletManager(),
		transactionManager: NewTransactionManager(),
		refundManager:      NewRefundManager(),
	}
}
func (s *Swiggy) addMoneyToWallet(userId string, amount float64) {
	s.walletManager.Add(userId, amount)
}
func (s *Swiggy) makePayment(userId string, amount float64) {
	s.walletManager.Deduct(userId, amount)
}
func (s *Swiggy) getTransactionHistory(userId string) {
	fmt.Println(s.transactionManager.GetAll(userId))
}
func (s *Swiggy) requestRefund(userId string) {
	s.refundManager.Refund(userId, "some_transaction_id", s.transactionManager)
}

func main() {
	fmt.Println("Welcome to swiggy Wallet")
	mySwiggy := NewSwiggyService()
	mySwiggy.addMoneyToWallet("user1", 100.0)
}
