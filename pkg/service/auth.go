package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"BST/pkg/repository"

	"BST/models"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type AuthService struct {
	repo repository.Users
}

func NewAuthService(repo repository.Users) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) GenerateToken(ctx context.Context, u models.User) (string, error) {
	if !u.Validate() {
		return "", models.ErrBadRequest
	}
	u.Password = generatePasswordHash(u.Password)
	userID, err := a.repo.AuthenticateUser(ctx, u)
	if err != nil {
		if err == models.ErrNoRows {
			return "", models.ErrUnauthorized
		}
		logrus.Println("GenerateToken a.db.AuthenticateUser", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(jwtString string) (int, error) {
	token, err := jwt.ParseWithClaims(jwtString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*models.JWTClaims)
	if !ok {
		return 0, fmt.Errorf("token claims are not of type *JWTClaims")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return 0, fmt.Errorf("token expired")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
