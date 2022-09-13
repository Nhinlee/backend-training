package db

import (
	"context"

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

func CreateRandomUser() (User, CreateUserParams, error) {
	randomUser := GetRandomUser()

	arg := CreateUserParams{
		FirstName: randomUser.FirstName,
		LastName:  randomUser.LastName,
		Email:     randomUser.Email,
		Password:  randomUser.Password,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	return user, arg, err
}
