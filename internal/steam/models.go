package steam

type gernerateSignedURLPayload struct{
	Id int`json:"id" validate:"required"`
	Email string `json:"email" validate:"required"`
}
type videoServiceInterface interface{

}

