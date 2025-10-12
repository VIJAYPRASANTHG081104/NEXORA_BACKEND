package users

import (
	"fmt"
	"net/http"
	"nexora_backend/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type Handler struct{
	store UserStoreInterface
}

func CreateUserHandler(store UserStoreInterface) *Handler{
	return &Handler{
		store: store,
	}
}

func (h* Handler) RegisterRoutes(router *gin.Engine){
	router.POST("/login", h.handleLogin)
	router.POST("/register", h.handleRegister)
}

func (h *Handler) handleLogin(c *gin.Context) {
	var payload LoginPayloadStruct

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Payload validation error: %v", errors)})
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	
	if err := utils.ComparePassword(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.CreateToken(user.Username, user.Id, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}


func(h *Handler) handleRegister(c *gin.Context){
	var payload RegisterPayloadStruct;
	
	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(http.StatusBadRequest,fmt.Errorf("invalid json payload"))
		return;
	}

	if err:= utils.Validate.Struct(payload);err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest,fmt.Errorf("playload Validation Error: %v",errors))
		return;
	}
	user,err := h.store.GetUserByEmail(payload.Email)

	if err != nil{
		c.JSON(http.StatusInternalServerError,err);
	}

	if user != nil {
		c.JSON(http.StatusBadRequest,fmt.Errorf("user already exist"));
	}

	payload.Password, err = utils.Encrypt(payload.Password)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,fmt.Errorf("failed to encrypt password"))
	}

	err = h.store.CreateUser(&payload);

	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"user created successfully",
	})
}

// Learn concreate type and type assertion in go
// interface type
// Name variable  differace private public
// interface rules
// migrations
// 