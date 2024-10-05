package admin

import (
	"context"
	"database/sql"
	"fmt"

	"myproject/pkg/model"
)

type Repository interface {
	Register(ctx context.Context, request model.AdminRegisterRequest) error

	Login(ctx context.Context, email string) (model.AdminRegisterRequest, error)
}

type repository struct {
	sql *sql.DB
}

func NewRepository(sqlDB *sql.DB) Repository {
	return &repository{
		sql: sqlDB,
	}
}

func (r *repository) Register(ctx context.Context, request model.AdminRegisterRequest) error {
	fmt.Println("this is in the repository Register")
	query := `INSERT INTO admin (name, gst, email, password,phone) VALUES ($1, $2, $3, $4,$5)`
	_, err := r.sql.ExecContext(ctx, query, request.Name, request.GST, request.Email, request.Password, request.Phone)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %w", err)
	}

	return nil
}

func (r *repository) Login(ctx context.Context, email string) (model.AdminRegisterRequest, error) {
	fmt.Println("theee !!!!!!!!!!!  LLLLoginnnnnn  ", email)
	query := `SELECT name, gst, email, password FROM admin WHERE email = $1`
	fmt.Println(`SELECT name, gst, email, password FROM admin WHERE email =  = 'adithyanunni258@gmail.com' ;`)

	var user model.AdminRegisterRequest
	err := r.sql.QueryRowContext(ctx, query, email).Scan(&user.Name, &user.GST, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.AdminRegisterRequest{}, nil
		}
		return model.AdminRegisterRequest{}, fmt.Errorf("failed to find user by email: %w", err)
	}
	fmt.Println("the data !!!! ", user)

	return user, nil
}
