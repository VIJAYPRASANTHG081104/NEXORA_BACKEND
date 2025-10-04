package users

import (
	"fmt"
	"net/http"
	"nexora_backend/types"
	"nexora_backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)


type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler{
	return &Handler{
		store: store,
	}
}

func (h* Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login",h.handleLogin);
	router.HandleFunc("/register",h.handleRegister);
}

func (h*Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	
}

func(h *Handler) handleRegister(w http.ResponseWriter, r*http.Request){
	var payload types.RegisterPayload;
	
	if err := utils.ParseJSON(r,&payload); err != nil{
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("Invalid json payload"))
		return;
	}

	if err:= utils.Validate.Struct(payload);err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w,http.StatusBadRequest,fmt.Errorf("playload Validation Error: %v",errors))
		return;
	}
}

// Learn concreate type and type assertion in go
