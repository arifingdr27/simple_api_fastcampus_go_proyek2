package memberships

import "gorm.io/gorm"

type User struct {
	Email        string `gorm:"uniqueIndex;not null" json:"email"`
	Username     string `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`
	CreatedBy    string `gorm:"not null" json:"created_by"`
	UpdatedBy    string `gorm:"not null" json:"updated_by"`
	gorm.Model
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
