package vendor

import (
	"context"
	"fmt"
	services "myproject/pkg/client"
	"myproject/pkg/config"
	"myproject/pkg/model"


	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(ctx context.Context, request model.VendorRegisterRequest) error
	Login(ctx context.Context, request model.VendorLoginRequest) error

	OtpLogin(ctx context.Context, request model.VendorOtp) error

	VerifyOtp(ctx context.Context, email string)

}

type service struct {
	repo     Repository
	Config   config.Config
	services services.Services
}

func NewService(repo Repository, services services.Services) Service {
	return &service{
		repo:     repo,
		services: services,
	}
}

func (s *service) VerifyOtp(ctx context.Context, email string) {
	s.repo.VerifyOtp(ctx, email)

}



func (s *service) Register(ctx context.Context, request model.VendorRegisterRequest) error {
	var err error
	if request.Name == "" || request.Email == "" || request.Password == "" || request.Phone == "" {
		fmt.Println("this is in the service error value missing")
		err = fmt.Errorf("missing values")
		return err
	}
	if !isValidEmail(request.Email) {
		fmt.Println("this is in the service error invalid email")
		err = fmt.Errorf("invalid email")
		return err
	}
	if !isValidPhoneNumber(request.Phone) {
		fmt.Println("this is in the service error invalid phone number")
		err = fmt.Errorf("invalid phone number")
		return err
	}
	fmt.Println("this is the dataaa ", request.Email)
	existingUser, err := s.repo.Login(ctx, request.Email)
	fmt.Println("there may be a user", existingUser)
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("this is in the service error checking existing user")
		err = fmt.Errorf("failed to check existing user: %w", err)
		return err
	}
	if existingUser.Email != "" {
		fmt.Println("this is in the service user already exists")
		err = fmt.Errorf("user already exists")
		return err
	}
	fmt.Println("this is in the service Register", request.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("this is in the service error hashing password")
		err = fmt.Errorf("failed to hash password: %w", err)
		return err
	}
	request.Password = string(hashedPassword)
	fmt.Println("this is in the service Register", request.Password)
	return s.repo.Register(ctx, request)
}

func (s *service) Login(ctx context.Context, request model.VendorLoginRequest) error {
	fmt.Println("this is in the service Login", request.Password)
	var err error
	if request.Email == "" || request.Password == "" {
		fmt.Println("this is in the service error value missing")
		err = fmt.Errorf("missing values")
		return err
	}
	storedUser, err := s.repo.Login(ctx, request.Email)
	fmt.Println("thisss is the dataaa ", storedUser)
	if err != nil {
		fmt.Println("this is in the service user not found")
		return fmt.Errorf("user not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(request.Password)); err != nil {
		fmt.Println("this is in the service incorrect password")
		return fmt.Errorf("incorrect password: %w", err)
	}

	return nil

	//return s.repo.Login(ctx, request)
}
func (s *service) OtpLogin(ctx context.Context, request model.VendorOtp) error {
	fmt.Println("this is in the service Login", request.Otp)
	var err error
	if request.Email == "" || request.Otp == "" {
		fmt.Println("this is in the service error value missing")
		err = fmt.Errorf("missing values")
		return err
	}
	// storedUser, err := s.repo.Login(ctx, request.Email)
	// if err != nil {
	// 	fmt.Println("this is in the service user not found")
	// 	return fmt.Errorf("user not found: %w", err)
	// }

	return nil

	//return s.repo.Login(ctx, request)
}

