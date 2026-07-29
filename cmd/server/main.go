package main
import(
	"github.com/AbhinanKumar/smart-dispatch/internal/router"
)
func main(){
	r := router.SetupRouter()
	r.Run(":8000")
}