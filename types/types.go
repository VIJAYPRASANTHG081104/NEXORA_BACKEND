package types


type UserStore interface{

}

type RegisterPayload struct{
	Username string `json:"username" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}