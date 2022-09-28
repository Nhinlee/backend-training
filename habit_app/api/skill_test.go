package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "habits.com/habit/db/mock"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/utils"
)

func TestGetSkillAPI(t *testing.T) {
	skill := randomSkill()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	arg := db.ListSkillsByUserParams{
		UserID: skill.UserID,
		Limit:  5,
		Offset: 0,
	}

	// Build stubs
	store.
		EXPECT().
		ListSkillsByUser(gomock.Any(), gomock.Eq(arg)).
		Times(1).Return([]db.Skill{skill}, nil)

	// Start server to send request
	server := NewTestServer(t, store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/skills?user_id=%d&page_id=%d&page_size=%d", skill.UserID, 1, 5)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)

	// Check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

func randomSkill() db.Skill {
	return db.Skill{
		SkillID: utils.RandomInt(1, 10000),
		UserID:  utils.RandomInt(1, 10000),
		Title:   utils.RandomString(20),
	}
}
