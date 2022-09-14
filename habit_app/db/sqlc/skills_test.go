package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSkill(t *testing.T) {
	CreateRandomSkill(t)
}

func TestUpdateSkill(t *testing.T) {
	oldSkill := CreateRandomSkill(t)

	arg := UpdateSkillParams{
		SkillID: oldSkill.SkillID,
		Title:   oldSkill.Title + " edited",
	}

	newSkill, err := testQueries.UpdateSkill(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, newSkill)
	require.Equal(t, arg.Title, newSkill.Title)
}
