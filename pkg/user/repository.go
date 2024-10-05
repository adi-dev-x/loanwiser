package user

import (
	"context"
	"database/sql"
	"fmt"

	"strings"

	"myproject/pkg/model"
)

// ListWish
type Repository interface {
	Login(ctx context.Context, email string) (model.UserRegisterRequest, error)
	Register(ctx context.Context, request model.UserRegisterRequest) (string, error)
	UpdateUser(ctx context.Context, query string, args []interface{}) error
	VerifyOtp(ctx context.Context, email string)
}

type repository struct {
	sql *sql.DB
}

func NewRepository(sqlDB *sql.DB) Repository {
	return &repository{
		sql: sqlDB,
	}
}
func (r *repository) VerifyOtp(ctx context.Context, email string) {
	query := `
	UPDATE users
	SET verification =true
	WHERE email = $1
	`

	_, err := r.sql.ExecContext(ctx, query, email)

	if err != nil {
		fmt.Errorf("failed to execute update query: %w", err)
	}

}

func (r *repository) Register(ctx context.Context, request model.UserRegisterRequest) (string, error) {
	fmt.Println("this is in the repository Register")
	var id string
	query := `INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4) Returning id`
	err := r.sql.QueryRowContext(ctx, query, request.FirstName, request.LastName, request.Email, request.Password).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to execute insert query: %w", err)
	}

	return id, nil
}

func (r *repository) Login(ctx context.Context, email string) (model.UserRegisterRequest, error) {
	fmt.Println("theee !!!!!!!!!!!  LLLLoginnnnnn  ", email)
	query := `SELECT firstname, lastname, email, password FROM users WHERE email = $1 AND verification=true`
	fmt.Println(`SELECT firstname, lastname, email, password FROM users WHERE email = 'adithyanunni258@gmail.com' ;`)

	var user model.UserRegisterRequest
	err := r.sql.QueryRowContext(ctx, query, email).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.UserRegisterRequest{}, nil
		}
		return model.UserRegisterRequest{}, fmt.Errorf("failed to find user by email: %w", err)
	}
	fmt.Println("the data !!!! ", user)

	return user, nil
}
func (r *repository) UpdateUser(ctx context.Context, query string, args []interface{}) error {
	queryWithParams := query
	for _, arg := range args {
		queryWithParams = strings.Replace(queryWithParams, "?", fmt.Sprintf("'%v'", arg), 1)
	}
	fmt.Println("Executing update with query:", queryWithParams)
	fmt.Println("Arguments:", args)
	fmt.Println("Executing update for email:", args[len(args)-1]) // Email is the last argument
	_, err := r.sql.ExecContext(ctx, queryWithParams)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}
	return nil
}
