package memberships

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"proyek3-catalog-music/internal/models/memberships"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_Login(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockService := NewMockservice(mockController)
	tests := []struct {
		name       string
		mockFn     func()
		statusCode int
		wantErr    bool
		resultBody memberships.LoginResponse
	}{
		{
			name: "sukses login controller",
			mockFn: func() {
				mockService.EXPECT().Login(memberships.LoginRequest{
					Email:    "testing123@gmail.com",
					Password: "testing123",
				}).Return("testing123", nil)
			},
			wantErr:    false,
			statusCode: 201,
			resultBody: memberships.LoginResponse{
				AccessToken: "testing123",
			},
		},
		{
			name: "gagal login controller",
			mockFn: func() {
				mockService.EXPECT().Login(memberships.LoginRequest{
					Email:    "testing123@gmail.com",
					Password: "testing123",
				}).Return("", assert.AnError)
			},
			wantErr:    true,
			statusCode: 400,
			resultBody: memberships.LoginResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.Default()
			h := &Handler{
				Engine:  api,
				service: mockService,
			}
			h.RegisterRoutes()
			responseNewRecorder := httptest.NewRecorder()
			endpoint := `/memberships/login`
			requestModel := memberships.LoginRequest{
				Email:    "testing123@gmail.com",
				Password: "testing123",
			}
			reqBody, err := json.Marshal(requestModel)
			assert.NoError(t, err)
			body := bytes.NewReader(reqBody)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(responseNewRecorder, req)
			fmt.Println(responseNewRecorder.Body)
			assert.Equal(t, tt.statusCode, responseNewRecorder.Code)
			if !tt.wantErr {
				res := responseNewRecorder.Result()
				defer res.Body.Close()

				response := memberships.LoginRequest{}
				err := json.Unmarshal(responseNewRecorder.Body.Bytes(), &response)
				assert.NoError(t, err)
			}
		})
	}
}
