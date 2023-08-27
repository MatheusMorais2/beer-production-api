package http

import (
	"beer-production-api/bootstrap"

	"github.com/labstack/echo/v4"
)

type Server struct {
	app *bootstrap.App
	echo *echo.Echo
}

func NewServer(app *bootstrap.App) *Server {
	e := echo.New()
	return &Server{ app: app, echo: e}
}
// @title Beer Production API
// @version 1.0
// @description Beer production API

// @contact.name Matheus Morais

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func (s *Server) Start() {
	//NewController aqui
	//Rotas e metodos aqui
	s.echo.Start(":8080")
}