package user

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Email string
	Wohnort string
	Oblast string
	School string
	Ort int
	OrtMat int
	OrtPhys int
}

func (u User) FirstName() string{
	return u.FirstName
}

func (u User) LastName() string{
	return u.LastName
}

func (u User) PhoneNumber() string{
	return u.PhoneNumber
}

func (u User) Email() string{
	return u.Email
}

func (u User) Wohnort() string{
	return u.Wohnort
}

func (u User) Oblast() string{
	return u.Oblast
}

func (u User) School() string{
	return u.School
}

func (u User) Ort() int{
	return u.Ort
}

func (u User) OrtMat() int{
	return u.OrtMat
}

func (u User) OrtPhys() int{
	return u.OrtPhys
}

