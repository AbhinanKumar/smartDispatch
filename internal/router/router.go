package router
import(
	"github.com/AbhinanKumar/smart-dispatch/internal/handler"
	"github.com/gin-gonic/gin"
)
func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine{
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.POST("/register", authHandler.Register)
	return r
}