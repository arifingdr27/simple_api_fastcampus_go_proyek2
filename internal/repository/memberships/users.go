package memberships

import (
	"errors"

	"proyek3-catalog-music/internal/models/memberships"
)

func (r *repository) CreateUser(user *memberships.User) error {
	if user == nil {
		return errors.New("user already exists")
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUser(email, username string, id int) (*memberships.User, error) {
	var user memberships.User
	if err := r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
