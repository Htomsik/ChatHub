package Server

import (
	"Chat/internal/app/model/chat"
	"Chat/internal/app/model/constants"
	"Chat/internal/app/store/serverStore/memoryStore"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestServer_ChatUseMiddleware check user validation
func TestServer_ChatUseMiddleware(t *testing.T) {

	// Arrange
	store := memoryStore.New()
	srv := newServer(store)

	// Create new cookie for response
	dummyHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	})

	cases := []struct {
		name         string
		model        chat.User
		expectedCode int
	}{
		{
			name:         "NoData",
			model:        chat.User{Name: ""},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "OK",
			model:        chat.User{Name: "TWO"},
			expectedCode: http.StatusOK,
		},
		{
			name:         "Less than 2",
			model:        chat.User{Name: "1"},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "More than 20",
			model:        chat.User{Name: "1111111111111111111111111111111111111111111111111111"},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "Less than 20",
			model:        chat.User{Name: "1111111111111111111111111111111111111111111111111111"},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "Invalid symbols",
			model:        chat.User{Name: "TWO@"},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	// Act
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/chat/1?%v=%v", constants.QueryValueUser, testCase.model.Name)

			request, _ := http.NewRequest(http.MethodGet, url, nil)

			srv.chatUserMiddleWare(dummyHandler).ServeHTTP(recorder, request)

			// Assert
			assert.Equal(t, testCase.expectedCode, recorder.Code)
		})
	}
}
