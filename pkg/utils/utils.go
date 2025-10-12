package utils

import (
	"fmt"
	"time"

	"nexora_backend/config"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()


func Encrypt(password string) (string,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil	
}

func ComparePassword(hashedPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func CreateToken(username,id,email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
         "email": email,
         "id":id,
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString([]byte(config.ENVS.JWTSecret))
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
   tokenString = tokenString[7:] // Remove "Bearer " prefix
   fmt.Println(config.ENVS.JWTSecret)
   fmt.Println(tokenString)
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     return []byte(config.ENVS.JWTSecret), nil
   })
  
   fmt.Println("token Valid",err)
   if err != nil {
      return err
   }
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}