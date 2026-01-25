package services

import (
	
	"time"

	"campus-connect-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string)(string, error){
	bytes, err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err
}

func CheckPassword(hash, password string) bool{
	return  bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))==nil
	
}

func GenerateJWT(userID string ,role string)(string, error){
	claims:=jwt.MapClaims{
		"user_id":userID,
		"role":role,
		"exp":time.Now().Add(24* time.Hour).Unix(),
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}