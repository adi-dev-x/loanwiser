package vendor

import (
	"context"
	"database/sql"
	"fmt"

	"myproject/pkg/model"
)

type Repository interface {
	Register(ctx context.Context, request model.VendorRegisterRequest) error

	Login(ctx context.Context, email string) (model.VendorRegisterRequest, error)

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
	UPDATE vendor
	SET verification =true
	WHERE email = $1
	`

	_, err := r.sql.ExecContext(ctx, query, email)

	if err != nil {
		fmt.Errorf("failed to execute update query: %w", err)
	}

}

func (r *repository) Register(ctx context.Context, request model.VendorRegisterRequest) error {
	fmt.Println("this is in the repository Register")
	query := `INSERT INTO vendor (name, gst, email, password,phone) VALUES ($1, $2, $3, $4,$5)`
	_, err := r.sql.ExecContext(ctx, query, request.Name, request.GST, request.Email, request.Password, request.Phone)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %w", err)
	}

	return nil
}

func (r *repository) Login(ctx context.Context, email string) (model.VendorRegisterRequest, error) {
	fmt.Println("Attempting to login with email:", email)

	// SQL query to fetch user details based on email
	query := `SELECT name, gst, email, password FROM vendor WHERE email = $1 AND verification=true`
	fmt.Printf("Executing query: %s\n", query)

	var user model.VendorRegisterRequest

	// Execute the query and scan the result into the user struct
	err := r.sql.QueryRowContext(ctx, query, email).Scan(&user.Name, &user.GST, &user.Email, &user.Password)
	if err != nil {
		// Check if no rows were returned
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the provided email.")
			return model.VendorRegisterRequest{}, nil
		}
		// For other types of errors, wrap and return the error
		return model.VendorRegisterRequest{}, fmt.Errorf("failed to find user by email: %w", err)
	}

	fmt.Println("User data retrieved:", user)

	return user, nil
}
