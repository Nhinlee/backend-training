package db

import (
	"habits.com/habit/utils"
)

func GetRandomUser() User {
	firstName := "test " + utils.RandomString(10)
	lastName := "test " + utils.RandomString(10)
	email := "test+" + utils.RandomString(10) + "@gmail.com"
	password := utils.RandomString(10)

	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
