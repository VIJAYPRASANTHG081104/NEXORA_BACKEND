package users

import (
	"database/sql"
	"fmt"
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

	if err := s.db.QueryRow("SELECT id,email FROM users WHERE id = $1",email).Scan(&user.Id, &user.Email); err != nil{
		if err == sql.ErrNoRows{
			return nil,nil
		}
		return  nil, fmt.Errorf("failed to fetch user emails")
	}
	return user,nil
}