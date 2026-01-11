package main

import (
	"math/rand"
	"strconv"
)

type User struct {
	id                      string
	notificationPreferences []ChannelType
}

type UserService struct {
	users map[string]*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

func (s *UserService) Create() *User {
	newUser := &User{
		id:                      "user" + strconv.Itoa(rand.Intn(1000)),
		notificationPreferences: make([]ChannelType, 0),
	}
	s.users[newUser.id] = newUser
	return newUser
}

func (s *UserService) Get(userId string) *User {
	return s.users[userId]
}

func (s *UserService) AddPreference(u *User, t ChannelType) {
	u.notificationPreferences = append(u.notificationPreferences, t)
}

func (s *UserService) GetAll() map[string]*User {
	return s.users
}
