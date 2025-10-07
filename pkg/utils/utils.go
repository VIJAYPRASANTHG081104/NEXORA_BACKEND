package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nexora_backend/config"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func ParseJSON(r *http.Request,value any) error{
	if r.Body == nil{
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(value);
}

func WriteJSON(w http.ResponseWriter, status int,value any) error{
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(value);
}

func WriteError(w http.ResponseWriter, status int, err error){
	WriteJSON(w,status,map[string]string{"error":err.Error()})
}

func Encrypt(password string) (string,error){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil	
}

func Comparepassword(hashedPassword, password string) error{
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

    tokenString, err := token.SignedString(config.ENVS.JWTSecret)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
   tokenString = tokenString[7:] // Remove "Bearer " prefix
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return config.ENVS, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}