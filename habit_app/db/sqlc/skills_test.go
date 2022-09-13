package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"habits.com/habit/utils"
)

func createRandomSkill(t *testing.T) Skill {
	user, _, _ := CreateRandomUser()

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

func TestCreateSkill(t *testing.T) {
	createRandomSkill(t)
}

func TestUpdateSkill(t *testing.T) {
	oldSkill := createRandomSkill(t)

	arg := UpdateSkillParams{
		SkillID: oldSkill.SkillID,
		Title:   oldSkill.Title + " edited",
	}

	newSkill, err := testQueries.UpdateSkill(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, newSkill)
	require.Equal(t, arg.Title, newSkill.Title)
}
