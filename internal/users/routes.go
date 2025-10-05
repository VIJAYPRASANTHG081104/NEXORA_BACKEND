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
	
}

func(h *Handler) handleRegister(w http.ResponseWriter, r*http.Request){
	fmt.Println("register");
	var payload RegisterPayloadStruct;
	
	if err := utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Invalid json payload"))
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

	

}
// Learn concreate type and type assertion in go
// interface type
// Name variable  differace private public
// interface rules
// migrations
// 