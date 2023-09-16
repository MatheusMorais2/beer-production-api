package auth

import (
	"encoding/json"
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

func GenerateUserToken(email string) (string, error) {
	rawToken, err := GenerateJWT()
	if err != nil {
		return "", fmt.Errorf("Error generating jwt token: %v", err)
	}

	claimedToken, err := ClaimJWT(rawToken, email)
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

func ClaimJWT(token *jwt.Token,  email string) (*jwt.Token, error) {
	claims := token.Claims.(jwt.MapClaims)
	expirationTime := time.Now().Add(10 * time.Minute)
	claims["exp"] = expirationTime
	claims["authorized"] = true
	claims["sub"] = email
	claims["iss"] = os.Getenv("JWT_ISSUER")


	return token, nil
}

func SignJWT(token *jwt.Token) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	fmt.Println(string(secretKey))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("e agr porra")
		return "", fmt.Errorf("Error signing token: %+v", err)
	}

	return tokenString, nil
}

func GetAuthTokenData(token string) (*jwt.Token, error) {
	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	VerifySignature(token)
	myMap := make(map[string]interface{})
	tokenData, err := jwt.Parse(
		token,
		func(t *jwt.Token) (interface{}, error) {
			claimsMap := t.Claims.(jwt.MapClaims)
			myMap["exp"] = claimsMap["exp"]
			myMap["sub"] = claimsMap["sub"]
			myMap["iss"] = claimsMap["iss"]
			myMap["authorized"] = claimsMap["authorized"]
			fmt.Printf("myMap: %+v\n", myMap)
			fmt.Printf("token dentro do keyFunc: %+v\n", t)
			return key, nil
		})
		
	fmt.Println(tokenData)
	fmt.Println(myMap)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}


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

func areValidClaims(claims jwt.Claims) bool {
/* 	fmt.Printf("claims.Valid(): %+v", claims)
	if err := claims.Valid() != nil; err {
		fmt.Println("Entrou no claims valid")
		return false
	} */
	fmt.Println("entrou aqui ó")
	claimsMap := claims.(jwt.MapClaims)

	if claimsMap["authorized"] != true {
		fmt.Println("Entrou no authorized false")
		return false
	}
	exp, ok := claimsMap["exp"]


	fmt.Printf("`ok: %+v\n`", ok)
	fmt.Printf("`exp: %+v\n`", exp)
	var expired bool
	switch exp.(type) {
	case float64:
		fmt.Println("O tipo do exp é: float64")
	case json.Number:
		fmt.Println("O tipo do exp é: json")
	case int64:
		fmt.Println("O tipo do exp é: int64")
	default:
		fmt.Println("Could not get expiration time type")
		return false
	}

	timeNow := time.Now().Unix()
	fmt.Printf("timeNow := int64(time.Now() %+v\n", timeNow)
	expired = claimsMap.VerifyExpiresAt(timeNow, true)
	if expired {
		fmt.Println("Entrou no expired")
		return false
	}

	//

	matchIssuer := claimsMap.VerifyIssuer(os.Getenv("JWT_ISSUER"), true)
	if !matchIssuer {
		fmt.Println("Entrou no issuer")
		return false
	}

	return true
}

/* func isExpired(exp interface{}) bool {
	switch exp.(type) {
	case float64:
		timeNow := time.Now().Unix()
		timeNow > exp
	case json.Number:
		fmt.Println("O tipo do exp é: json")
	case int64:
		fmt.Println("O tipo do exp é: int64")
	default:
		fmt.Println("Could not get expiration time type")
		return false
	}
} */

func VerifySignature(tokenString string) bool {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	parts := strings.Split(tokenString, ".")
    if len(parts) != 3 {
        fmt.Println("Token JWT inválido")
        return false
    }
	fmt.Printf("\nparts: %+v\n", parts)

    // Decodificar as partes base64
    header, err := base64.RawURLEncoding.DecodeString(parts[0])
    if err != nil {
        fmt.Println("Erro ao decodificar o cabeçalho:", err)
        return false
    }
	fmt.Printf("\nheader: %+v\n", header)

    payload, err := base64.RawURLEncoding.DecodeString(parts[1])
    if err != nil {
        fmt.Println("Erro ao decodificar o payload:", err)
        return false
    }
	fmt.Printf("\npayload: %+v\n", payload)

    signature, err := base64.RawURLEncoding.DecodeString(parts[2])
    if err != nil {
        fmt.Println("Erro ao decodificar a assinatura:", err)
        return false
    }
	fmt.Printf("\nsignature: %+v\n", signature)

    // Verificar a assinatura
    hasher := hmac.New(sha256.New, []byte(secretKey))
    hasher.Write([]byte(parts[0] + "." + parts[1]))
    expectedSignature := hasher.Sum(nil)
	fmt.Printf("\nexpectedSignature: %+v\n", expectedSignature)

    if hmac.Equal(signature, expectedSignature) {
        fmt.Println("A verificação da assinatura foi bem-sucedida")
		return true
    } else {
        fmt.Println("A verificação da assinatura falhou")
		return false
    }

}