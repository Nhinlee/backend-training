package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
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

func CreateRandomUser(t *testing.T) User {
	randomUser := GetRandomUser()

	arg := CreateUserParams{
		FirstName: randomUser.FirstName,
		LastName:  randomUser.LastName,
		Email:     randomUser.Email,
		Password:  randomUser.Password,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.UserID)

	return user
}

func CreateRandomSkill(t *testing.T) Skill {
	user := CreateRandomUser(t)

	title := "test skill " + utils.RandomString(20)

	arg := CreateSkillParams{
		UserID: user.UserID,
		Title:  title,
	}

	skill, err := testQueries.CreateSkill(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, skill)
	require.Equal(t, arg.Title, skill.Title)
	require.Equal(t, arg.UserID, skill.UserID)

	return skill
}

func CreateRandomHabit(t *testing.T, user *User) Habit {

	title := "habit test " + utils.RandomString(20)
	targetConsecutiveDays := 21

	arg := CreateHabitParams{
		UserID:                user.UserID,
		Title:                 title,
		TargetConsecutiveDays: int32(targetConsecutiveDays),
	}

	habit, err := testQueries.CreateHabit(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, habit)
	require.Equal(t, arg.Title, habit.Title)
	require.Equal(t, arg.SkillID, habit.SkillID)
	require.Equal(t, arg.UserID, user.UserID)

	return habit
}
