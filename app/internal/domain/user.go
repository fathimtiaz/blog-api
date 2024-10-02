package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           int
	Name         string
	Email        string
	PasswordHash string

	Timestamp
}

func NewUser(name, email, password string) (User, error) {
	if hashedPass, err := hashPassword(password); err != nil {
		return User{}, err
	} else {
		return User{
			Name:         name,
			Email:        email,
			PasswordHash: hashedPass,
			Timestamp:    Timestamp{CreatedAt: time.Now()},
		}, nil
	}
}

func hashPassword(password string) (string, error) {
	if byteHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return "", err
	} else {
		return string(byteHash), nil
	}
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

func (u *User) GenerateToken(secret string, hourExpired int) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*time.Duration(hourExpired)).Unix(), 0)),
		Subject:   u.Email,
	}).SignedString([]byte(secret))
}
