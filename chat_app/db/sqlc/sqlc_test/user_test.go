package sqlc_test

import (
	"context"
	"testing"
	core "v1/core"
	db "v1/db/sqlc"
	utils "v1/utils"

	"github.com/stretchr/testify/require"
)

func GetRandomUser() db.User {
	userId := utils.RandomID()
	firstName := "test " + utils.RandomString(10)
	lastName := "test " + utils.RandomString(10)
	email := "test+" + utils.RandomString(10) + "@gmail.com"
	password := utils.RandomString(10)
	hashedPassword, _ := core.HashPassword(password)

	return db.User{
		UserID:         userId,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		HashedPassword: hashedPassword,
	}
}

func CreateRandomUser(t *testing.T) db.User {
	randomUser := GetRandomUser()

	arg := db.CreateUserParams{
		UserID:         randomUser.UserID,
		FirstName:      randomUser.FirstName,
		LastName:       randomUser.LastName,
		Email:          randomUser.Email,
		HashedPassword: randomUser.HashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.NotZero(t, user.UserID)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1, user2)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	arg := db.UpdateUserInfoParams{
		UserID:    user1.UserID,
		FirstName: user1.FirstName + "updated",
		LastName:  user1.LastName + "updated",
	}

	user2, err := testQueries.UpdateUserInfo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.UserID, user1.UserID)
	require.Equal(t, user2.FirstName, arg.FirstName)
	require.Equal(t, user2.LastName, arg.LastName)
}
