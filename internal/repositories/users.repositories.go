package repositories

import (
	"context"
	"errors"
	"main/internal/models"
	"main/pkg"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepositories struct{}

var UserRepo *UserRepositories

func NewUserRepository() {
	UserRepo = &UserRepositories{}
}

func (u *UserRepositories) AddNewUser(c context.Context, email, password string) (pgconn.CommandTag, error) {
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	values := []any{email, password}
	cmd, err := pkg.Database.Exec(c, query, values...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	return cmd, nil
}

func (u *UserRepositories) IsUserExist(c context.Context, email string) (bool, error) {
    query := "SELECT 1 FROM users WHERE email=$1"
    row := pkg.Database.QueryRow(c, query, email)

    var exists int
    err := row.Scan(&exists)

    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return false, nil
        }
        return false, err
    }
    return true, nil
}

func (u *UserRepositories) LoginUser(ctx context.Context, email, password string) (int, error) {
	query := "SELECT id FROM users WHERE email=$1 AND password=$2"
	var userID int

	err := pkg.Database.QueryRow(ctx, query, email, password).Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, errors.New("user not exist")
		}
		return 0, err
	}

	return userID, nil
}

func (u *UserRepositories) FindUserByEmail(ctx context.Context, email string) (models.Users, error) {
    query := "SELECT email, password FROM users WHERE email = $1"

    var user models.Users
    err := pkg.Database.QueryRow(ctx, query, email).Scan(&user.Email, &user.Password)
    if err != nil {
        return models.Users{}, err
    }
    return user, nil
}
