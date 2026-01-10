package main

type Wallet struct {
	id      string
	balance float64
	userId  string
}

type WalletManager struct {
	wallets map[string]*Wallet
}

func NewWalletManager() *WalletManager {
	return &WalletManager{
		wallets: make(map[string]*Wallet),
	}
}

func (mgr *WalletManager) getWallet(userId string) *Wallet {
	if _, ok := mgr.wallets[userId]; ok {
		return mgr.wallets[userId]
	} else {
		mgr.wallets[userId] = &Wallet{
			id:      "newId",
			balance: 0.0,
			userId:  userId,
		}
		return mgr.wallets[userId]
	}
}

func (mgr *WalletManager) Get(userId string) float64 {
	myWallet := mgr.getWallet(userId)
	return myWallet.balance
}

func (mgr *WalletManager) Add(userId string, amount float64) {
	myWallet := mgr.getWallet(userId)
	myWallet.balance += amount
}

func (mgr *WalletManager) Deduct(userId string, amount float64) {
	myWallet := mgr.getWallet(userId)
	myWallet.balance -= amount
}
