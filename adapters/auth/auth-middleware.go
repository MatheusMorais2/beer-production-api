package auth

import (
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
		
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest,`Could not get token data: `+ err.Error())
		}

		userId, valid := areValidClaims(tokenData.Claims)
		if (valid != nil) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
		}

		c.Set("userId", userId)

		next(c)
		return nil
	}
}


