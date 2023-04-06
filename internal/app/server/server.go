package server

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/web/docs"
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/handlers"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/store"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/gofiber/template/html"
)

type server struct {
	router   *fiber.App
	logger   logging.Logger
	pgStore  store.PostgresStore
	mgStore  store.MongoStore
	config   *config.Config
	handlers *handlers.Handlers
}

func newServer(pgstore store.PostgresStore, mgstore store.MongoStore, config *config.Config, log logging.Logger) *server {
	//ViewsLayout: "shared/headers/main_header"
	engine := html.New("./templates", ".html")
	fmt.Println(&engine)
	s := &server{
		router:   fiber.New(fiber.Config{Views: engine, ViewsLayout: "shared/main_layout", ServerHeader: "software engineering course api", AppName: "Api v1.0.1"}),
		logger:   log,
		pgStore:  pgstore,
		mgStore:  mgstore,
		config:   config,
		handlers: handlers.NewHandlers(pgstore, mgstore, log),
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	//api := s.router.Group("/api")
	//api.Use(logger.New())
	// localhost:4000/user/register

	s.router.Get("/swagger/*", swagger.HandlerDefault)

	///////// USER GROUP ///////////////
	////////////////////////////////////
	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	user.Post("/register", s.handlers.Register())
	user.Post("/login", s.handlers.Login())
	user.Post("/check", s.handlers.CheckJWT())
	//////////////////////////////////////

	pages := s.router.Group("/")
	pages.Static("/public", "./public")
	pages.Get("auth", s.handlers.AuthPage())
	pages.Get("main", s.handlers.MainPage())

}