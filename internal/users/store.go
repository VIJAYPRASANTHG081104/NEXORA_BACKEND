package users

import (
	"database/sql"
	"fmt"
	"strings"
)

type Store struct {
	db* sql.DB
}

func NewUserStore(db* sql.DB) *Store{
	return &Store{
		db:db,
	}
}

func (s* Store) GetUserByEmail(email string) (*UserStruct,error){
	user :=new(UserStruct) //new keyword create memory return pointer of that type

	if err := s.db.QueryRow("SELECT id,email,username, password FROM users WHERE email = $1",email).Scan(&user.Id, &user.Email,&user.Username,&user.Password); err != nil{
		if err == sql.ErrNoRows{
			return nil,nil
		}
		return  nil, fmt.Errorf("failed to fetch user emails")
	}
	return user,nil
}

func (s *Store) CreateUser(payload *RegisterPayloadStruct) (error){
	_, err := s.db.Exec("INSERT INTO users (username,email,password) VALUES ($1,$2,$3)", strings.Split(payload.Email, "@")[0],payload.Email,payload.Password)
	if err != nil{
		return fmt.Errorf("failed to create user: %v",err)
	}
	return nil
}		
