package handler

import(
	"github.com/AbhinanKumar/smart-dispatch/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/AbhinanKumar/smart-dispatch/internal/dto"
)

type AuthHandler struct{
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler{
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context){
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil{			//pointer -> Gin needs to fill the struct.
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return 
	}

	res, err := h.authService.Register(req)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(c *gin.Context){
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid json body",
		})
		return
	}
	res, err := h.authService.Login(req)
	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Profile(c *gin.Context){
	userID, _ := c.Get("userID")
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"email": email,
	})
}