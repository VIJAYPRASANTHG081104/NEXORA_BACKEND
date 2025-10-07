package users

import (
	"fmt"
	"net/http"
	"nexora_backend/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type Handler struct{
	store UserStoreInterface
}

func NewHandler(store UserStoreInterface) *Handler{
	return &Handler{
		store: store,
	}
}

func (h* Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin);
	router.HandleFunc("/register",h.handleRegister).Methods("POST");
}

func (h*Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayloadStruct;
	if err := utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid json payload"))
		return;
	}

	if err:= utils.Validate.Struct(payload);err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("playload Validation Error: %v",errors))
		return;
	}

	user, err := h.store.GetUserByEmail(payload.Email);
	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,err);
		return;
	}

	if user == nil {
		utils.WriteError(w,http.StatusUnauthorized,fmt.Errorf("invalid email or password"));
		return;
	}

	if err := utils.Comparepassword(user.Password,payload.Password); err != nil{
		utils.WriteError(w,http.StatusUnauthorized,fmt.Errorf("invalid email or password"));
		return;
	}

	jwt,err := utils.CreateToken(user.Username,user.Id,user.Email);

	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,fmt.Errorf("failed to generate token"));
		return;
	}
	utils.WriteJSON(w,http.StatusOK,map[string]string{"message":"login successful","token":jwt,});
}

func(h *Handler) handleRegister(w http.ResponseWriter, r*http.Request){
	var payload RegisterPayloadStruct;
	fmt.Println(r.Body)
	
	if err := utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("invalid json payload"))
		return;
	}

	if err:= utils.Validate.Struct(payload);err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("playload Validation Error: %v",errors))
		return;
	}
	user,err := h.store.GetUserByEmail(payload.Email)

	if err != nil{
		utils.WriteError(w,http.StatusInternalServerError,err);
	}

	if user != nil {
		utils.WriteJSON(w,http.StatusBadRequest,fmt.Errorf("user already exist"));
	}

	payload.Password, err = utils.Encrypt(payload.Password)
	
	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,fmt.Errorf("failed to encrypt password"))
	}

	err = h.store.CreateUser(&payload);

	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
		return
	}
	utils.WriteJSON(w,http.StatusCreated,map[string]string{"message":"user created successfully"})
}

// Learn concreate type and type assertion in go
// interface type
// Name variable  differace private public
// interface rules
// migrations
// 