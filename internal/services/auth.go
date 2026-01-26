package services

import (
	"errors"
	
	"time"

	"campus-connect-backend/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaims struct{
	UserID  string  `json:"user_id"`
	Role    string  `json:"role"`
	jwt.RegisteredClaims
}

func HashPassword(password string)(string, error){
	bytes, err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err
}

func CheckPassword(hash, password string) bool{
	return  bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))==nil
	
}

func GenerateJWT(userID string ,role string)(string, error){
	claims:=JWTClaims{
		UserID: userID,
		Role:role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func ValidateJWT(tokenStr string)(*JWTClaims, error){
	token , err:=jwt.ParseWithClaims(tokenStr,&JWTClaims{},func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
	})
	if err !=nil{
		return nil,err
	}

	claims, ok :=token.Claims.(*JWTClaims)

	if !ok || !token.Valid{
		return nil, errors.New("invalid token")
	}
	return claims,nil
}