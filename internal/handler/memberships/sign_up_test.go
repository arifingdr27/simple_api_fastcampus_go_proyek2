package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"proyek3-catalog-music/internal/models/memberships"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mocksrvc := NewMockservice(ctrlMock)

	tests := []struct {
		name               string
		mockFn             func()
		requestBody        interface{}
		expectedStatusCode int
	}{
		{
			name: "berhasil",
			mockFn: func() {
				mocksrvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "testing@gmail.com",
					Username: "testing",
					Password: "testing_pass", // ✅ Sesuaikan dengan request body
				}).Return(nil)
			},
			requestBody: memberships.SignUpRequest{
				Email:    "testing@gmail.com",
				Username: "testing",
				Password: "testing_pass",
			},
			expectedStatusCode: 201,
		},
		{
			name: "gagal - service error",
			mockFn: func() {
				mocksrvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "testing@gmail.com",
					Username: "testing",
					Password: "testing_pass",
				}).Return(errors.New("user already exists"))
			},
			requestBody: memberships.SignUpRequest{
				Email:    "testing@gmail.com",
				Username: "testing",
				Password: "testing_pass",
			},
			expectedStatusCode: 500, // ✅ Sesuai dengan implementasi handler yang mengembalikan 500
		},
		{
			name: "gagal - invalid JSON",
			mockFn: func() {
				// No mock expectation because service won't be called
			},
			requestBody:        "invalid json",
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()

			api := gin.Default()
			h := &Handler{
				Engine:  api,
				service: mocksrvc,
			}
			h.RegisterRoutes()
			w := httptest.NewRecorder()
			endpoint := `/memberships/signup`

			val, err := json.Marshal(tt.requestBody)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
