package auth

import (
	"awcoding.com/back/domain/users"
	"awcoding.com/back/infrastructure/config"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type AuthService interface {
	SignIn(login string, password string) (*Auth, error)
	GetByToken(token string) (*users.User, error)
}

type Service struct {
	userService users.UserService
	cfg         *config.Config
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewService(userService users.UserService, cfg *config.Config) *Service {
	return &Service{userService: userService, cfg: cfg}
}

func (s *Service) SignIn(login string, password string) (*Auth, error) {
	hashedPassword := s.generatePasswordHash(password)
	log.Println(hashedPassword)
	user, err := s.userService.GetByEmailPassword(login, hashedPassword)
	if err != nil {
		return nil, errors.New("login and password incorrect")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &Auth{Token: token, User: user}, nil
}

func (s *Service) GetByToken(token string) (*users.User, error) {
	id, err := s.parseToken(token)
	if err != nil {
		return nil, err
	}
	user, err := s.userService.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) generateToken(user *users.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(s.cfg.Secret.TokenKey))
}

func (s *Service) parseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.Secret.TokenKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *Service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.Secret.Salt)))
}
