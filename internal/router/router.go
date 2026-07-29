package router
import(
	"github.com/AbhinanKumar/smart-dispatch/internal/handler"
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/health", handler.Health)
	return r
}