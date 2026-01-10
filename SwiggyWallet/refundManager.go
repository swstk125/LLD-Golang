package main

type RefundManager struct{}

func NewRefundManager() *RefundManager {
	return &RefundManager{}
}

func (mgr *RefundManager) ValidateRefund(userId string, transactionId string, transactionMgr *TransactionManager) bool {
	// validationlogic
	return true
}

func (mgr *RefundManager) Refund(userId string, transactionId string, transactionMgr *TransactionManager) {
	if mgr.ValidateRefund(userId, transactionId, transactionMgr) {
		amount := transactionMgr.Get(userId, transactionId).amount
		transactionMgr.Add(userId, amount, "success", "refund")
	}
}
