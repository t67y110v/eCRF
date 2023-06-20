package server

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/web/docs"
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/handlers"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/middlewares"
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
				ServerHeader: "eCRF API",
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
		AllowOrigins:     s.config.FrontPORT,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	s.router.Get("/swagger/*", swagger.HandlerDefault)

	///////// USER GROUP ///////////////
	////////////////////////////////////
	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Get("/new", s.handlers.NewUser())
	user.Post("/register", s.handlers.Register())
	user.Post("/login", s.handlers.Login(), s.pages.AuthPage())

	user.Post("/update", s.handlers.Update())
	//////////////////////////////////////

	pages := s.router.Group("/")
	pages.Static("/public", "./public")
	pages.Get("auth", s.pages.AuthPage())
	pages.Use(middlewares.CheckJWT())
	pages.Get("main/filter:filter", s.pages.MainPage())
	pages.Get("protocol/:id/:number", s.pages.ProtocolPage())
	pages.Get("journal", s.pages.JournalPage())

	protocol := s.router.Group("/protocols")
	protocol.Use(logger.New())
	protocol.Use(middlewares.CheckJWT())
	protocol.Post("/save", s.handlers.SaveProtocol())
	protocol.Post("/add", s.handlers.AddProtocol())
	protocol.Post("/delete", s.handlers.DeleteProtocol())

	adminPanel := s.router.Group("/admin")
	adminPanel.Use(middlewares.CheckJWT())
	adminPanel.Get("/panel", s.pages.AdminPage())

	center := s.router.Group("/center")
	center.Use(logger.New())
	center.Use(middlewares.CheckJWT())
	center.Post("/new", s.handlers.AddNewCenter())
	center.Post("/update", s.handlers.UpdateCenter())

	subject := s.router.Group("/subject")
	subject.Use(middlewares.CheckJWT())
	subject.Post("/new", s.handlers.NewSubject())
	screening := subject.Group("/screening")
	screening.Post("/informaionconsent", s.handlers.InformaionConsentSubject())
	screening.Post("/demography", s.handlers.DemographySubject())
	screening.Post("/anthropometry", s.handlers.AnthropometrySubject())
	screening.Post("/inclusioncriteria", s.handlers.InclusionCriteriaSubject())
	screening.Post("/exclusioncriteria", s.handlers.ExclusionСriteriaSubject())
	screening.Post("/completion", s.handlers.CompletionOfScreening())
	screening.Post("/updatecolor", s.handlers.UpdateColor())
	screening.Post("/updatecolorwithcomment", s.handlers.UpdateColorWithComment())
	screening.Post("/updatefield", s.handlers.UpdateFieldValue())
	errors := s.router.Group("/error")
	errors.Get("/", s.pages.ErrorPage())

}
