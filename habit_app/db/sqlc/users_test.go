package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	user, arg, err := CreateRandomUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.UserID)
}

func TestGetUser(t *testing.T) {
	user1, _, err := CreateRandomUser()
	user2, err := testQueries.GetUser(context.Background(), user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1, user2)
}

func TestUpdateUser(t *testing.T) {
	user1, _, err := CreateRandomUser()

	arg := UpdateUserInfoParams{
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
