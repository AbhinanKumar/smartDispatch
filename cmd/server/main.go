package main
import(
	"log"
	"github.com/AbhinanKumar/smart-dispatch/internal/config"
	"github.com/AbhinanKumar/smart-dispatch/internal/router"
	"github.com/joho/godotenv"
)
func main(){
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}
	_, err := config.NewDatabase()
	if err != nil{
		log.Fatal(err)
	}

	r := router.SetupRouter()
	log.Println("Server running on :8000")
	r.Run(":8000")
}