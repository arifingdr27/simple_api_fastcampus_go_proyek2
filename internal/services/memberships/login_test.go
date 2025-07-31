package memberships

import (
	"proyek3-catalog-music/internal/configs"
	"proyek3-catalog-music/internal/models/memberships"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func Test_service_Login(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockRepository := NewMockrepository(mockController)
	type args struct {
		req memberships.LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "sukses service login",
			args: args{
				req: memberships.LoginRequest{
					Email:    "testing_arifin@gmail.com",
					Password: "testing",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(args.req.Password), bcrypt.DefaultCost)
				assert.NoError(t, err)
				mockRepository.EXPECT().GetUser(args.req.Email, "", 0).Return(&memberships.User{
					Email:        args.req.Email,
					PasswordHash: string(hashedPassword),
					Username:     "testing",
				}, nil)
			},
		},
		{
			name: "gagal service login",
			args: args{
				req: memberships.LoginRequest{
					Email:    "testing_arifin@gmail.com",
					Password: "testing",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepository.EXPECT().GetUser(args.req.Email, "", 0).Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJwt: "testingJWT",
					},
				},
				repository: mockRepository,
			}
			got, err := s.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got)
			} else {
				assert.Empty(t, got)
			}
		})
	}
}
