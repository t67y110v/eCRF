package server

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/web/docs"
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/handlers"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/pages"
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
	pages    *pages.Pages
}

func newServer(
	pgstore store.PostgresStore,
	mgstore store.MongoStore,
	config *config.Config,
	log logging.Logger,
) *server {

	engine := html.New("./templates", ".html")
	var st state
	st.store = pgstore
	var c counter
	engine.AddFuncMap(template.FuncMap{
		"setCounter": c.SetCounter,
		"incCounter": c.IncCounter,
		"set":        st.Set,
		"inc":        st.Inc,
		"com":        st.Com,
		"role":       st.Role,
		"center":     st.Center,
	})
	s := &server{
		router: fiber.New(
			fiber.Config{
				Views:        engine,
				ViewsLayout:  "shared/main_layout",
				ServerHeader: "software engineering course api",
				AppName:      "Api v1.0.1",
			}),
		logger:   log,
		pgStore:  pgstore,
		mgStore:  mgstore,
		config:   config,
		handlers: handlers.NewHandlers(pgstore, mgstore, log),
		pages:    pages.NewPages(pgstore, mgstore, log),
	}
	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	s.router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
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
	user.Post("/login", s.handlers.Login(), s.pages.AuthPage())
	user.Post("/check", s.handlers.CheckJWT())
	user.Post("/update", s.handlers.Update())
	//////////////////////////////////////

	pages := s.router.Group("/")
	pages.Static("/public", "./public")
	pages.Get("auth", s.pages.AuthPage())
	pages.Get("main/filter:filter", s.pages.MainPage())
	pages.Get("protocol/:id/:number", s.pages.ProtocolPage())
	pages.Get("protocol/edit/:id", s.pages.ProtocolEdit())
	pages.Get("protocol/", s.pages.ProtocolNew())

	protocol := s.router.Group("/protocols")
	protocol.Use(logger.New())
	protocol.Post("/save", s.handlers.SaveProtocol())
	protocol.Post("/add", s.handlers.AddProtocol())

	adminPanel := s.router.Group("/admin")
	adminPanel.Get("/panel", s.pages.AdminPage())
	adminPanel.Get("/update/:email", s.pages.UpdatePage())

	center := s.router.Group("/center")
	center.Use(logger.New())
	center.Post("/new", s.handlers.AddNewCenter())
	center.Post("/update", s.handlers.UpdateCenter())

	subject := s.router.Group("/subject")
	subject.Post("/new", s.handlers.NewSubject())

	errors := s.router.Group("/error")
	errors.Get("/", s.pages.ErrorPage())

}
