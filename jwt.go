package rice

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token with the given username and expiration time
func GenerateToken(claim MyCustomClaims) (string, error) {

	// Define a secret key for signing the token
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", jwt.ErrInvalidKey
	}

	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken parses the JWT token and returns the claims
func ParseToken(tokenString string) (*MyCustomClaims, error) {
	// Define a secret key for signing the token
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, jwt.ErrInvalidKey
	}

	// Parse the token and validate it
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrTokenInvalidClaims
	}
}

// ValidateToken validates the JWT token and returns true if valid, false otherwise
func ValidateToken(tokenString string) (bool, error) {
	// Define a secret key for signing the token
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return false, jwt.ErrInvalidKey
	}

	// Parse the token and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
