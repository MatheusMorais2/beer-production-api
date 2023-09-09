package auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware (next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		fmt.Println("chegou no middleware")
		req := c.Request()
		authorizationHeader := req.Header.Get("Authorization")
		token, err := GetAuthTokenFromHeader(authorizationHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Could not get authorization header")
		}
		fmt.Printf("token: %+v\n", token)
		next(c)
		return nil
	}
}


