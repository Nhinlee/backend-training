package sqlc_test

import (
	"context"
	"testing"
	db "v1/db/sqlc"

	"github.com/stretchr/testify/require"
)

func TestCreateConversationUser(t *testing.T) {
	conversation := CreateRandomConversation(t)

	user1 := CreateRandomUser(t)
	user2 := CreateRandomUser(t)

	c_user1, err := testQueries.CreateConversationUser(context.Background(), db.CreateConversationUserParams{
		UserID:         user1.UserID,
		ConversationID: conversation.ConversationID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, c_user1)

	c_user2, err := testQueries.CreateConversationUser(context.Background(), db.CreateConversationUserParams{
		UserID:         user2.UserID,
		ConversationID: conversation.ConversationID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, c_user2)

}
