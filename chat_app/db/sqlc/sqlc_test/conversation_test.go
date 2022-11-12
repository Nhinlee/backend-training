package sqlc_test

import (
	"context"
	"database/sql"
	"testing"
	db "v1/db/sqlc"
	"v1/utils"

	"github.com/stretchr/testify/require"
)

func getRandomConversation() db.Conversation {
	return db.Conversation{
		ConversationID: utils.RandomID(),
		ConversationName: sql.NullString{
			String: "Conversation " + utils.RandomString(10),
			Valid:  true,
		},
	}
}

func createRandomConversation(t *testing.T) db.Conversation {
	randomConversation := getRandomConversation()

	args := db.CreateConversationParams{
		ConversationID:   randomConversation.ConversationID,
		ConversationName: randomConversation.ConversationName,
	}

	conversation, err := testQueries.CreateConversation(context.Background(), args)

	require.NotEmpty(t, conversation)
	require.NoError(t, err)

	require.Equal(t, args.ConversationID, conversation.ConversationID)
	require.Equal(t, args.ConversationName, conversation.ConversationName)

	return conversation
}

func TestCreateConversation(t *testing.T) {
	createRandomConversation(t)
}

func TestGetConversation(t *testing.T) {
	randomConversation := createRandomConversation(t)

	conversation, err := testQueries.GetConversationById(context.Background(), randomConversation.ConversationID)

	require.NoError(t, err)
	require.NotEmpty(t, conversation)

	require.Equal(t, randomConversation, conversation)
}

func TestUpdateConversation(t *testing.T) {
	randomConversation := createRandomConversation(t)

	args := db.UpdateConversationInfoParams{
		ConversationID: randomConversation.ConversationID,
		ConversationName: sql.NullString{
			String: randomConversation.ConversationName.String + "Updated",
			Valid:  true,
		},
	}

	updatedConversation, err := testQueries.UpdateConversationInfo(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updatedConversation)

	require.Equal(t, args.ConversationID, updatedConversation.ConversationID)
	require.Equal(t, args.ConversationName, updatedConversation.ConversationName)
}

func TestDeleteConversation(t *testing.T) {
	randomConversation := createRandomConversation(t)

	err := testQueries.DeleteConversation(context.Background(), randomConversation.ConversationID)

	require.NoError(t, err)
}
