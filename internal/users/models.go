package users

type RegisterPayloadStruct struct{
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserStoreInterface interface{
	GetUserByEmail(email string) (*UserStruct,error)
}

type UserStruct struct{
	Id string
	Username string 
	Email string 
	Password string
}