package application

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/wodm8/wodm8-core/initializers"
	"github.com/wodm8/wodm8-core/internal/domain"
	"github.com/wodm8/wodm8-core/internal/members"
	"github.com/wodm8/wodm8-core/internal/users"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsersService struct {
	usersRepository users.UsersRepository
}

func NewUsersService(usersRepository users.UsersRepository, memberRepository members.MemberRepository) UsersService {
	return UsersService{
		usersRepository: usersRepository,
	}
}

func (u UsersService) CreateUser(userReq domain.CreateUserRequest) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)

	if err != nil {
		fmt.Errorf("failed to hash password: %v", err)
		return err
	}

	user, err := users.NewUsers(userReq.Id, userReq.FirstName, userReq.LastName, userReq.Email, string(hash))
	fmt.Printf("new user: %v\n", user)

	if err != nil {
		fmt.Errorf("error creating user: %v", err)
		return err
	}

	err = u.usersRepository.Save(user)
	if err != nil {
		fmt.Errorf("error saving user: %v", err)
		return err
	}

	return nil
}

func (u UsersService) Login(loginReq domain.LoginRequest) (string, error) {
	var cfg = initializers.Cfg

	user, err := u.usersRepository.First(loginReq.Email)
	if err != nil {
		fmt.Errorf("error getting user: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		fmt.Errorf("error getting user: %v", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JwtSecret))

	if err != nil {
		fmt.Errorf("failed to create token: %v", err)
		return "", err
	}

	return tokenString, nil
}
