package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func NewClaims(user_id int, role string) *Claims {
	return &Claims{
		UserID: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)), // jwt aktif selama 5 menit
		},
	}
}

func (c *Claims) GenerateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("secret not provided")
	}
	// 1. buat token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 2. tanda tangan token
	return token.SignedString([]byte(jwtSecret))
}

func (c *Claims) VerifyToken(token string) error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("secret not provided")
	}
	parsedToken, err := jwt.ParseWithClaims(token, c, func(t *jwt.Token) (interface{}, error) {
		// fungsi callback yang digunakan oleh ParseWithClaims untuk mengambil secret
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return errors.New("expired token")
	}
	return nil
}
