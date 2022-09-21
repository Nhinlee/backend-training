package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createHabitLogByUser(t *testing.T, user *User, habit *Habit, dateTime int64) HabitLog {
	arg := CreateHabitLogParams{
		HabitID:  habit.HabitID,
		UserID:   user.UserID,
		DateTime: dateTime,
	}
	habitLog, err := testQueries.CreateHabitLog(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, habitLog)
	require.Equal(t, habitLog.UserID, user.UserID)
	require.Equal(t, habitLog.HabitID, habit.HabitID)

	return habitLog
}

func TestCreateHabitLog(t *testing.T) {
	user := CreateRandomUser(t)
	habit := CreateRandomHabit(t, &user)

	createHabitLogByUser(t, &user, &habit, 0)
}

func TestGetHabitLogByUser(t *testing.T) {
	user := CreateRandomUser(t)
	habit1 := CreateRandomHabit(t, &user)
	habit2 := CreateRandomHabit(t, &user)
	habitIds := map[int64]bool{
		habit1.HabitID: true,
		habit2.HabitID: true,
	}

	createHabitLogByUser(t, &user, &habit1, 0)
	createHabitLogByUser(t, &user, &habit2, 0)

	habitLogs, err := testQueries.GetHabitLogsByUser(context.Background(), user.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, habitLogs)
	require.Equal(t, len(habitLogs), 2)
	for _, habitLog := range habitLogs {
		require.Equal(t, habitIds[habitLog.HabitID], true)
	}
}

func TestGetLatestHabitLogByUserIsEmpty(t *testing.T) {
	user := CreateRandomUser(t)

	habitLogs, err := testQueries.GetLatestHabitLogByUser(context.Background(), GetLatestHabitLogByUserParams{
		UserID:  user.UserID,
		HabitID: -1,
	})
	require.NoError(t, err)
	require.Equal(t, habitLogs, []HabitLog([]HabitLog{}))
}
