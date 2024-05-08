package jwt

import (
	"GoLinkTree/custom"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	// "github.com/joho/godotenv"
)

func CreateToken(username string) (string, error){
	// err := godotenv.Load()
    // if err != nil {
    //     return "", err
    // }

    SECRET_KEY := os.Getenv("SECRET_KEY")
	
	if SECRET_KEY == "" {
		err := custom.MyError("SECRET KEY not found in env")
		return "", err
	}

	// Create a new token object, specifying signing method and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}