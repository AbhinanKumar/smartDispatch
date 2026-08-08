package main
import(
	"log"
	"github.com/AbhinanKumar/smart-dispatch/internal/config"
	"github.com/AbhinanKumar/smart-dispatch/internal/router"
	"github.com/joho/godotenv"
	"github.com/AbhinanKumar/smart-dispatch/internal/repository"
	"github.com/AbhinanKumar/smart-dispatch/internal/service"
	"github.com/AbhinanKumar/smart-dispatch/internal/handler"
)
func main(){
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}
	db, err := config.NewDatabase()
	if err != nil{
		log.Fatal(err)
	}

	userRepo :=  repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	customerRepo := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)


	r := router.SetupRouter(authHandler, customerHandler,)
	log.Println("Server running on :8000")
	r.Run(":8000")
}