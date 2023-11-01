package http

import (
	"beer-production-api/adapters/auth"
	"beer-production-api/bootstrap"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "beer-production-api/adapters/http/docs"
)

type Server struct {
	app *bootstrap.App
	echo *echo.Echo
}

func NewServer(app *bootstrap.App) *Server {
	e := echo.New()
	return &Server{app: app, echo: e}
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
	userController := NewUserController(s.app)
	breweryController := NewBreweryController(s.app)
	recipeController := NewRecipeController(s.app)
	batchController := NewBatchController(s.app)

	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	  }))

	s.echo.GET("/docs/*", echoSwagger.WrapHandler)
	s.echo.POST("/users", userController.CreateUser)
	s.echo.POST("/auth/login", userController.Login)

	s.echo.POST("/brewery", breweryController.CreateBrewery, auth.AuthMiddleware)

	s.echo.POST("/recipes", recipeController.CreateRecipe, auth.AuthMiddleware)

	s.echo.POST("/batches", batchController.CreateBatch, auth.AuthMiddleware)

	s.echo.POST("/start-batch", batchController.CreateBatchStep, auth.AuthMiddleware)

	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}