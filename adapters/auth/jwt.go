package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type jwtClaims struct {
    authorized bool
	exp time.Time
	iss string
	sub string
}

func GenerateUserToken(id string) (string, error) {
	rawToken, err := GenerateJWT()
	if err != nil {
		return "", fmt.Errorf("Error generating jwt token: %v", err)
	}

	claimedToken, err := ClaimJWT(rawToken, id)
	if err != nil {
		return "", fmt.Errorf("Error claiming jwt token: %v", err)
	}

	signedToken, err := SignJWT(claimedToken)
	if err != nil {
		return "", fmt.Errorf("Error signing jwt token: %v", err)
	}

	return signedToken, nil
}

func GenerateJWT() (*jwt.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	if token == nil {
		return nil, fmt.Errorf("Error generating token")
	}

	return token, nil
}

func ClaimJWT(token *jwt.Token, id string) (*jwt.Token, error) {
	claims := token.Claims.(jwt.MapClaims)
	expirationTime := time.Now().Add(10 * time.Minute)
	claims["exp"] = expirationTime
	claims["authorized"] = true
	claims["sub"] = id
	claims["iss"] = os.Getenv("JWT_ISSUER")


	return token, nil
}

func SignJWT(token *jwt.Token) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Error signing token: %+v", err)
	}

	return tokenString, nil
}

func GetAuthTokenData(token string) (*jwt.Token, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	verifySignature(token)

	myMap := make(map[string]interface{})

	tokenData, _ := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			claimsMap := t.Claims.(jwt.MapClaims)

			myMap["exp"] = claimsMap["exp"]
			myMap["sub"] = claimsMap["sub"]
			myMap["iss"] = claimsMap["iss"]
			myMap["authorized"] = claimsMap["authorized"]

			return key, nil
		})

	return tokenData, nil
}

func GetAuthTokenFromHeader(s string) (string, error) {
	stringArray := strings.Split(s, "Bearer ")

	if len(stringArray) != 2 {
		err := fmt.Errorf("Authorization header malformed")
		return "", err
	}

	if len(stringArray[1]) == 0 {
		err := fmt.Errorf("Authorization header malformed")
		return "", err
	}
	
	return stringArray[1], nil
}

func areValidClaims(claims jwt.Claims) (userId interface{}, err error) {
	claimsMap := claims.(jwt.MapClaims)

	if claimsMap["authorized"] != true {
		return nil, fmt.Errorf("Unathorized")
	}

	exp, ok := claimsMap["exp"]

	if !ok {
		return nil, fmt.Errorf("Unathorized")
	}

	if isExpired(exp) {
		return nil, fmt.Errorf("Expired token")
	}

	matchIssuer := claimsMap.VerifyIssuer(os.Getenv("JWT_ISSUER"), true)

	if !matchIssuer {
		return nil, fmt.Errorf("Wrong token issuer")
	}

	sub, ok := claimsMap["sub"]

	if !ok {
		return nil, fmt.Errorf("Unathorized")
	}

	if sub == "" {
		return nil, fmt.Errorf("Missing user id") 
	}

	return sub, nil
}

func isExpired(exp interface{}) bool {
	expiration, err := time.Parse("2006-01-02T15:04:05.999999999-07:00", fmt.Sprintf("%+v", exp))
	if err != nil {
		return true
	}
	
	timeNow := time.Now()

	isExpired := timeNow.After(expiration)
	if isExpired {
		return true
	}

	return false
}

func verifySignature(tokenString string) bool {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	parts := strings.Split(tokenString, ".")

    if len(parts) != 3 {
        return false
    }

    _, err := base64.RawURLEncoding.DecodeString(parts[0])
    if err != nil {
        return false
    }


    _, err = base64.RawURLEncoding.DecodeString(parts[1])
    if err != nil {
        return false
    }


    signature, err := base64.RawURLEncoding.DecodeString(parts[2])
    if err != nil {
        return false
    }

    hasher := hmac.New(sha256.New, []byte(secretKey))
    hasher.Write([]byte(parts[0] + "." + parts[1]))
    expectedSignature := hasher.Sum(nil)

    if hmac.Equal(signature, expectedSignature) {
		return true
    } else {
		return false
    }

}