package videos

import (
	"fmt"
	"net/http"
	"nexora_backend/pkg/utils"
	"os"

	"github.com/gin-gonic/gin"
)


type VideoServiceHandler struct {
	store videoServiceInterface
}

func CreateVideoServiceHandler(store videoServiceInterface) (*VideoServiceHandler ){
	return &VideoServiceHandler{
		store: store,
	}
}


func (h *VideoServiceHandler) VideoServiceRouter(rg *gin.RouterGroup) {
	rg.GET("/generateSignedURL", generateSignedURL)
	rg.PUT("/generateSignedURL", generateSignedURL)
}


func generateSignedURL(c *gin.Context)  {
	var payload gernerateSignedURLPayload;

	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(http.StatusBadRequest,fmt.Errorf("invalid json payload"))
		return;
	}
	signedURL, err := utils.GenerateV4PutObjectSignedURL(os.Getenv("GCP_VIDEO_BUCKET"),fmt.Sprintf("video/%d/%s",payload.Id,payload.Email),c.Request.Method);
	if err != nil {
		c.JSON(http.StatusInternalServerError,fmt.Errorf("unable to generate signed url: %v",err))
		return;
	}

	response := map[string]string{
		"signedURL": signedURL,
	}

	c.JSON(http.StatusOK,response);
}