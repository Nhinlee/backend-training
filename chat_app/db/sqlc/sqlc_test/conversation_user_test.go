package sqlc_test

import (
	"context"
	"testing"
	db "v1/db/sqlc"

	"github.com/stretchr/testify/require"
)

func CreateConversationUser(t *testing.T, conversationID string) db.ConversationUser {
	user := CreateRandomUser(t)

	c_user, err := testQueries.CreateConversationUser(context.Background(), db.CreateConversationUserParams{
		UserID:         user.UserID,
		ConversationID: conversationID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, c_user)

	return c_user
}

func TestCreateConversationUser(t *testing.T) {
	conversation := CreateRandomConversation(t)
	conversationID := conversation.ConversationID

	CreateConversationUser(t, conversationID)
	CreateConversationUser(t, conversationID)
	CreateConversationUser(t, conversationID)
	CreateConversationUser(t, conversationID)
}

func TestGetUserIDByConversationID(t *testing.T) {
	conversation := CreateRandomConversation(t)
	conversationID := conversation.ConversationID

	userIDs := []string{
		CreateConversationUser(t, conversationID).UserID,
		CreateConversationUser(t, conversationID).UserID,
		CreateConversationUser(t, conversationID).UserID,
		CreateConversationUser(t, conversationID).UserID,
	}

	actualUserIDs, err := testQueries.ListUserIdByConversationId(context.Background(), conversationID)
	require.NoError(t, err)
	require.NotEmpty(t, actualUserIDs)
	require.Equal(t, userIDs, actualUserIDs)
}
