package main

import "time"

type Status string
type TransactionType string

const (
	Success Status = "success"
	Failed  Status = "failed"
)

const (
	CREDIT TransactionType = "credit"
	DEBIT  TransactionType = "debit"
	REFUND TransactionType = "refund"
)

type transaction struct {
	id              string
	userId          string
	amount          float64
	status          Status
	transactionType TransactionType
	timestamp       time.Time
}

type TransactionManager struct {
	TransactionList map[string][]*transaction
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		TransactionList: make(map[string][]*transaction),
	}
}

func (mgr *TransactionManager) create(userId string, amount float64, transactionStatus Status, transactionType TransactionType) *transaction {
	return &transaction{
		id:              "newTransaction",
		userId:          userId,
		amount:          amount,
		status:          transactionStatus,
		transactionType: transactionType,
		timestamp:       time.Now(),
	}
}

func (mgr *TransactionManager) Add(userId string, amount float64, transactionStatus Status, transactionType TransactionType) {
	mgr.TransactionList[userId] = append(
		mgr.TransactionList[userId],
		mgr.create(userId, amount, transactionStatus, transactionType),
	)
}

func (mgr *TransactionManager) GetAll(userId string) []*transaction {
	return mgr.TransactionList[userId]
}

func (mgr *TransactionManager) Get(userId string, transactionId string) *transaction {
	for _, data := range mgr.TransactionList[userId] {
		if data.id == transactionId {
			return data
		}
	}
	return nil
}
