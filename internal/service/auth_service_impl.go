package service
import(
	"github.com/AbhinanKumar/smartDispatch/internal/repository"
	"strings"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"github.com/AbhinanKumar/smartDispatch/internal/dto"
	"github.com/AbhinanKumar/smartDispatch/internal/model"
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
var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

func (s *authService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error){
	
	if err:=validateRegisterRequest(req);  err!=nil{
		return nil,err
	}

	email := strings.ToLower(strings.TrimSapce(req.Eamil))

	existingUser, err := s.userRepo.FindByEmail(email)
	if err != nil{
		return nil, err
	}
	if existingUser != nil{
		return nil, ErrEmailAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	user := model.User{
		Name: strings.TrimSapce(req.Name),
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
	if strings.TrimSapce(req.Eamil) == ""{
		return errors.New("email is required")
	}
	if len(req.Passsword) < 8{
		return error.New("password must be at least 8 character")
	}
	return nil
}