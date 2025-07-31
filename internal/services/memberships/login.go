package memberships

import (
	"errors"
	"proyek3-catalog-music/internal/models/memberships"
	"proyek3-catalog-music/pkg/jwt"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) Login(req memberships.LoginRequest) (string, error) {
	user, err := s.repository.GetUser(req.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user data")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exists")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", err
	}

	accessToken, err := jwt.CreateToken(int64(user.ID), user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		log.Error().Err(err).Msg("failed to create jwt token")
		return "", err
	}

	return accessToken, nil
}
