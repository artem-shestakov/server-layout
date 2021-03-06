package auth

import (
	apperror "api/internal/apperrors"
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int
}

type AuthConfig struct {
	PasswordSalt string
	JWTKey       string
	JWTTTL       time.Duration
}

type AuthService struct {
	storage Authorization
	config  AuthConfig
}

func NewAuthService(storage *Storage, config *AuthConfig) *AuthService {
	return &AuthService{
		storage: storage,
		config:  *config,
	}
}

func (s *AuthService) CreateUser(user *CreateUser) (*User, error) {
	role, err := s.storage.GetRole("User")
	if err != nil {
		return nil, err
	}
	hashPass, err := s.hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashPass
	user.Role = role.Name
	return s.storage.CreateUser(user)
}

func (s *AuthService) GenerateJWT(email, password string) (string, error) {
	hashPassword, err := s.hashPassword(password)
	if err != nil {
		return "", err
	}
	user, err := s.storage.GetUserID(email, hashPassword)
	if err != nil {
		return "", err
	}

	JWTToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.config.JWTTTL * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return JWTToken.SignedString([]byte(s.config.JWTKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.config.JWTKey), nil
	})
	if err != nil {
		return 0, apperror.NewError(err, http.StatusInternalServerError, "Internal error", err.Error())
	}
	if !token.Valid {
		return 0, apperror.NewError(nil, http.StatusUnauthorized, "Token is not valid", "")
	}
	if claims, ok := token.Claims.(*tokenClaims); !ok {
		return 0, apperror.NewError(nil, http.StatusInternalServerError, "Internal error", "Token claims are not of type \"tokenClaims\"")
	} else {
		return claims.UserId, nil
	}
}

func (s *AuthService) hashPassword(password string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", apperror.NewError(err, http.StatusInternalServerError, "Internal server error", err.Error())
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(s.config.PasswordSalt))), nil
}
