package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaims interface {
    Valid() error
}

func GenerateUserToken(email string) (string, error) {
	rawToken, err := GenerateJWT()
	if err != nil {
		return "", err
	}

	claimedToken, err := ClaimJWT(rawToken, email)
	if err != nil {
		return "", err
	}

	signedToken, err := SignJWT(claimedToken)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateJWT() (*jwt.Token, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	if token == nil {
		fmt.Errorf("Error generating token")
	}

	return token, nil
}

func ClaimJWT(token *jwt.Token,  email string) (*jwt.Token, error) {
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["sub"] = email
	claims["iss"] = os.Getenv("JWT_issuer")

	return token, nil
}

func SignJWT(token *jwt.Token) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Errorf("Error signing token")
		return "", err
	}

	return tokenString, nil
}

func GetAuthTokenData(token string, key []byte) (*jwt.Token, error) {
	tokenData, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {return key, nil})
	if err != nil {
		err = fmt.Errorf("Error parsing token data: %+v", err)
		return nil, err
	}
	fmt.Printf("tokenData: %+v", tokenData)
	return tokenData, nil
}

func GetAuthTokenFromHeader(s string) (string, error) {
	stringArray := strings.Split(s, "Bearer ")
	if len(stringArray) != 2 {
		err := fmt.Errorf("Authorization header malformed")
		return "", err
	}
	return stringArray[1], nil
}