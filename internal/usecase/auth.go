package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

const (
	signingKey = "qrkjk#4#5FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int          `json:"user_id"`
	Roles  []dtos.Roles `json:"roles"`
}

func (u *usecase) CreateUserAsClient(input dtos.User) (int, error) {
	input.Password = generateHashPassword(input.Password)
	return u.repos.CreateUserAsClient(input)
}

func (u *usecase) GenerateToken(username, password string) (string, []dtos.Roles, error) {
	user, err := u.repos.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", nil, err
	}

	log.Printf("Here is user: %s", user.Username)

	var rolesHeaders []dtos.Roles
	roles, err := u.repos.GetRoles(user.UserID)
	log.Println(roles)
	if err != nil {
		return "", nil, err
	}

	for _, role := range roles {
		role_id, err := u.repos.GetRoleId(role, user.UserID)
		log.Printf("%s role and his id %d (log from generate token method\n)", role, role_id)
		if err != nil {
			return "", nil, err
		}
		rolesHeaders = append(rolesHeaders, dtos.Roles{RoleId: role_id, RoleName: role})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserID,
		rolesHeaders,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", nil, err
	}

	return tokenString, rolesHeaders, nil
}

func (u *usecase) ParseToken(tokenString string) (int, []dtos.Roles, error) {

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil, err
	}

	if !token.Valid {
		return 0, nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Roles, nil
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
