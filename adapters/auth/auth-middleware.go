package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware (next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		req := c.Request()

		authorizationHeader := req.Header.Get("Authorization")

		token, err := GetAuthTokenFromHeader(authorizationHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Could not get authorization header")
		}

		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header missing")
		}

		tokenData, err := GetAuthTokenData(token)
		fmt.Printf("tokenData que eles me dao : ", tokenData)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest,`Could not get token data: `+ err.Error())
		}

		valid := areValidClaims(tokenData.Claims)
		if (!valid) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
		}

		next(c)
		return nil
	}
}


