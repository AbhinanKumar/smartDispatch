package router
import(
	"github.com/AbhinanKumar/smart-dispatch/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/AbhinanKumar/smart-dispatch/internal/middleware"
)
func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine{
	r := gin.Default()
	r.GET("/health", handler.Health)
	r.POST("/register", authHandler.Register)

	r.POST("/login", authHandler.Login)

	protected := r.Group("/")
	protected.Use(midddleware.AuthMiddleware()) 
	
	protected.GET("/profile", authHandler.Profile)
	
	return r
}