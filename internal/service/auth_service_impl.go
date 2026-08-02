package service
import(
	"github.com/AbhinanKumar/smart-dispatch/internal/repository"
	"strings"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/AbhinanKumar/smart-dispatch/internal/dto"
	"github.com/AbhinanKumar/smart-dispatch/internal/model"
	appErrors "github.com/AbhinanKumar/smart-dispatch/internal/errors"
	"github.com/AbhinanKumar/smart-dispatch/internal/config"
)

type authService struct{
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService{
	return &authService{
		userRepo: userRepo,
	}
}

/*  Notice: userRepo repository.UserRepository
	Not: *repository.GormUserRepository

	Why? -> Because the service depends on an abstraction, not a concrete implementation.
	Tomorrow we can inject: Mock repository, Mongo repository, PostgreSQL repository without changing the service.
*/

// var (
// 	ErrEmailAlreadyExists = errors.New("email already exists")
// 	ErrInvalidCredentials = errors.New("invalid email or password")
// )

func (s *authService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error){
	
	if err:=validateRegisterRequest(req);  err!=nil{
		return nil,err
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))

	existingUser, err := s.userRepo.FindByemail(email)
	if err != nil{
		return nil, err
	}
	if existingUser != nil{
		return nil, appErrors.ErrEmailAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	user := model.User{
		Name: strings.TrimSpace(req.Name),
		Email: email,
		PasswordHash: string(hash),
	}

	if err = s.userRepo.Create(&user); err != nil{
		return nil, err
	}
	return &dto.RegisterResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}

func validateRegisterRequest(req dto.RegisterRequest) error{
	if strings.TrimSpace(req.Name) == ""{
		return errors.New("name is required")
	}
	if strings.TrimSpace(req.Email) == ""{
		return errors.New("email is required")
	}
	if len(req.Password) < 8{
		return errors.New("password must be at least 8 character")
	}
	return nil
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error){
	if err := validateLoginRequest(req); err != nil {
		return nil, err
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))

	user, err := s.userRepo.FindByemail(email); 
	if err != nil{
		return nil, err
	}
	if user == nil{
		return nil, appErrors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil{
		return nil, appErrors.ErrInvalidCredentials
	}

	token, err := config.GenerateJWT(user.ID, user.Email)
	if err != nil{
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func validateLoginRequest(req dto.LoginRequest) error{
	if strings.TrimSpace(req.Email) == ""{
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Password) == ""{
		return errors.New("password is required")
	}

	return nil
}