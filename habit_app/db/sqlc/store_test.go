package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"habits.com/habit/utils"
)

func TestCreateHabitAndSkill(t *testing.T) {
	store := NewStore(testDB)

	user := CreateRandomUser(t)

	errs := make(chan error)
	results := make(chan CreateHabitAndSkillResult)

	// run concurrent create skill and user at the same time
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			skillTitle := utils.RandomString(20)
			habitTitle := utils.RandomString(20)
			maxConsecutiveDays := int32(utils.RandomInt(20, 100))

			rs, err := store.CreateHabitAndSkill(context.Background(), CreateHabitAndSkillTxParams{
				UserID:             user.UserID,
				SkillTitle:         skillTitle,
				HabitTitle:         habitTitle,
				MaxConsecutiveDays: maxConsecutiveDays,
			})

			errs <- err
			results <- rs
		}()
	}

	for i := 0; i < n; i++ {
		// check error
		err := <-errs
		require.NoError(t, err)

		// check rs
		rs := <-results
		require.NotEmpty(t, rs)

		// check skill
		skill := rs.NewSkill
		require.NotEmpty(t, skill)

		// check habit
		habit := rs.NewHabit
		require.NotEmpty(t, habit)
	}
}
