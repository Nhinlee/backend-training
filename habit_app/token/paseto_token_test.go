package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"habits.com/habit/utils"
)

func TestPasetoFactory(t *testing.T) {
	factory, err := NewPasetoTokenFactory(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomString(20)
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := factory.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := factory.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.UserName)

	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpireToken(t *testing.T) {
	factory, err := NewPasetoTokenFactory(utils.RandomString(32))
	require.NoError(t, err)

	token, err := factory.CreateToken(utils.RandomString(20), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := factory.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpireToken.Error())
	require.Nil(t, payload)
}
