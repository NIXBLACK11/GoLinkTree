package jwt

import (
	"GoLinkTree/custom"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func AuthToken(username string, token string) (bool, error) {
	err := godotenv.Load()
	if err != nil {
		return false, err
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")

	if SECRET_KEY == "" {
		err := custom.MyError("SECRET KEY not found in env")
		return false, err
	}

	// Parse the JWT token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if !parsedToken.Valid {
		return false, nil
	}

	// Extract claims from the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("failed to parse claims")
	}

	// Extract username from claims
	claimedUsername, ok := claims["username"].(string)
	if !ok {
		return false, fmt.Errorf("invalid username claim")
	}

	// Compare the provided username with the username from the token
	if claimedUsername != username {
		return false, nil
	}

	// If all checks passed, the token is valid and corresponds to the provided username
	return true, nil
}
