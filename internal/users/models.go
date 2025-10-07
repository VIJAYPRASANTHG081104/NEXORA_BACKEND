package users

type RegisterPayloadStruct struct{
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserStoreInterface interface{
	GetUserByEmail(email string) (*UserStruct,error)
	CreateUser(payload *RegisterPayloadStruct) error
}

type UserStruct struct{
	Id string
	Username string 
	Email string 
	Password string
}

type LoginPayloadStruct struct{
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}