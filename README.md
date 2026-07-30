# smartDispatch
handling the request from customer portal and dispatching it to appropriate technician

# Install Dependencies
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv


Q1. Why don't we connect to PostgreSQL inside every repository?
Because creating database connections is expensive. We create one shared *gorm.DB instance and inject it into repositories. GORM manages a connection pool internally, so repositories reuse existing connections instead of opening new ones.

Q2. Why is *gorm.DB passed as a pointer?
Because it represents a shared database handle and connection pool. Copying it is unnecessary; all repositories should operate on the same underlying pool.

Q3. Why use AutoMigrate in development?
It keeps the schema synchronized with our models during development. In production, teams often prefer explicit migration tools for versioned and controlled schema changes.


If there were no pool, every request would do: 
Open TCP Connection -> Authenticate -> Execute Query -> Close Connection

With Connection Pooling:
main.go -> gorm.Open() -> Pool Created