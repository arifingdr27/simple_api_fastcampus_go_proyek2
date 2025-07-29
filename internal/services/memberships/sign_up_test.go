package memberships

import (
	"database/sql"
	"testing"

	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/internal/models/memberships"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success signup",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "testing@proyek2.com",
					Username: "testing",
					Password: "testing",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				// Mock GetUser mengembalikan error sql.ErrNoRows untuk user tidak ditemukan
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, 0).Return(nil, sql.ErrNoRows)
				// mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
				mockRepo.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(user *memberships.User) error {
					// Verify password hash is created correctly
					err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(args.request.Password))
					assert.NoError(t, err)
					return nil
				})
			},
		},
		{
			name: "failed signup - user already exists",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "existing@proyek2.com",
					Username: "existing",
					Password: "testing",
				},
			},
			mockFn: func(args args) {
				// Mengembalikan user yang sudah ada
				existingUser := &memberships.User{
					Email:        "existing@proyek2.com",
					Username:     "existing",
					PasswordHash: "hashedpassword",
				}
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, 0).Return(existingUser, nil)
			},
		},
		{
			name: "failed create user",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "testing@test.com",
					Username: "testing",
					Password: "testiing",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, 0).Return(nil, sql.ErrNoRows)
				// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
