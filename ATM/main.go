package main

import "fmt"

/*
Design aTM
	- user has a card, account
	- user can use card in the atm to access account
		- check / withdraw / deposit
	- atm needs to authorise user card via a pin
	- atm provides responses to all user actions

	entities
		- user
		- card
		- account -> AccountService
		- ATM
			- idle
			- cardInserted
			- checkBalance
			- withdraw
				- ejectCash
			- deposit
				- insertCash
			- eject
*/

type User struct {
	id string
}

type Card struct {
	number string
	userId string
}

type Account struct {
	id     string
	userId string
	pin    uint32
}

type AccountService struct {
	userAccounts map[string]*Account // userId -> account
}

func NewAccountService() *AccountService {
	return &AccountService{
		userAccounts: make(map[string]*Account),
	}
}

func (s *AccountService) CreateAccount(userId string, pin uint32) *Account {
	// check if user has already an account
	newAccount := &Account{
		id:     "account" + userId,
		userId: userId,
		pin:    pin,
	}

	s.userAccounts[userId] = newAccount

	return newAccount
}

func (s *AccountService) GetAccount(userId string, pin uint32) (*Account, error) {
	// check if user has an account
	if _, ok := s.userAccounts[userId]; !ok {
		return nil, fmt.Errorf("Account does not exist")
	}
	return s.userAccounts[userId], nil
}

type AtmState interface {
	insertCard(atm *AtmMachine, c *Card) error
	enterPin(atm *AtmMachine, pin uint32) error
	checkBalance(atm *AtmMachine) error
	withdraw(atm *AtmMachine, amt int) error
	deposit(atm *AtmMachine, amt int) error
	exitAtm(atm *AtmMachine) error
}

type IdleState struct{}
type CardInsertedState struct{}
type AuthenticatedState struct{}

func (i *IdleState) insertCard(atm *AtmMachine, c *Card) error {
	atm.card = c
	atm.SetState(&CardInsertedState{})
	return nil
}
func (i *IdleState) enterPin(atm *AtmMachine, pin uint32) error {
	return fmt.Errorf("Card not inserted")
}
func (i *IdleState) checkBalance(atm *AtmMachine) error      { return fmt.Errorf("Card not inserted") }
func (i *IdleState) withdraw(atm *AtmMachine, amt int) error { return fmt.Errorf("Card not inserted") }
func (i *IdleState) deposit(atm *AtmMachine, amt int) error  { return fmt.Errorf("Card not inserted") }
func (i *IdleState) exitAtm(atm *AtmMachine) error           { return fmt.Errorf("Card not inserted") }

func (i *CardInsertedState) insertCard(atm *AtmMachine, c *Card) error {
	return fmt.Errorf("Card already inserted")
}
func (i *CardInsertedState) enterPin(atm *AtmMachine, pin uint32) error {
	//validate pin
	return nil
}
func (i *CardInsertedState) checkBalance(atm *AtmMachine) error {
	return fmt.Errorf("PIN not entered")
}
func (i *CardInsertedState) withdraw(atm *AtmMachine, amt int) error {
	return fmt.Errorf("PIN not entered")
}
func (i *CardInsertedState) deposit(atm *AtmMachine, amt int) error {
	return fmt.Errorf("PIN not entered")
}
func (i *CardInsertedState) exitAtm(atm *AtmMachine) error {
	return nil
}

type AtmMachine struct {
	state      AtmState
	accService *AccountService
	card       *Card
}

func GetAtm(srv *AccountService) *AtmMachine {
	return &AtmMachine{}
}

func (atm *AtmMachine) SetState(s AtmState) {
	atm.state = s
}
