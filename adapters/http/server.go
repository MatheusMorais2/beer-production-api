package http

import (
	"beer-production-api/bootstrap"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"beer-production-api/adapters/auth"
	_ "beer-production-api/adapters/http/docs"
)

type Server struct {
	app *bootstrap.App
	echo *echo.Echo
	db *sql.DB
}

func NewServer(app *bootstrap.App, db *sql.DB) *Server {
	e := echo.New()
	return &Server{app: app, echo: e, db: db}
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
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	  }))

	s.echo.GET("/healthz", s.checkHealth)

	s.echo.GET("/docs/*", echoSwagger.WrapHandler)
	s.echo.POST("/users", userController.CreateUser)
	s.echo.POST("/auth/login", userController.Login)

	s.echo.POST("/brewery", breweryController.CreateBrewery, auth.AuthMiddleware)
	s.echo.GET("/brewery", breweryController.GetBreweriesByUserId, auth.AuthMiddleware)
	s.echo.POST("/brewery/invite", breweryController.InviteUser, auth.AuthMiddleware)

	s.echo.POST("/recipes", recipeController.CreateRecipe, auth.AuthMiddleware)
	s.echo.GET("/recipes/brewery/:brewery-id", recipeController.GetRecipesByBreweryId, auth.AuthMiddleware)
	s.echo.GET("/recipes/:brewery-id/steps/:recipe-id", recipeController.GetRecipeSteps, auth.AuthMiddleware)

	s.echo.POST("/batches", batchController.CreateBatch/* , auth.AuthMiddleware */)

	s.echo.POST("/start-batch", batchController.CreateBatchStep/* , auth.AuthMiddleware */)

	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

type DatabaseError struct {
	Message string `json:"message"`
}

func(s *Server) checkHealth(c echo.Context) error {
	err := s.db.Ping()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, DatabaseError{Message: "Database is not healthy: " + err.Error()})
	}
	
	return c.NoContent(200)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}