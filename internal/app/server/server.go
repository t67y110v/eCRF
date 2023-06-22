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
		AllowOrigins:     s.config.FrontURL,
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))
	s.router.Get("/swagger/*", swagger.HandlerDefault)

	///////// USER GROUP ///////////////
	////////////////////////////////////
	user := s.router.Group("/user")
	user.Use(logger.New())
	user.Get("/new", s.handlers.NewUser())
	user.Post("/login", s.handlers.UserLogin())
	user.Use(middlewares.CheckJWT())
	user.Post("/register", s.handlers.UserRegister())
	user.Patch("/update", s.handlers.UserUpdate())
	user.Delete("/delete", s.handlers.UserDelete())
	user.Get("/all", s.handlers.GetUsers())
	//////////////////////////////////////

	pages := s.router.Group("/")
	pages.Static("/public", "./public")
	pages.Get("auth", s.pages.AuthPage())
	pages.Get("main/filter:filter", s.pages.MainPage())
	pages.Get("protocol/:id/:number", s.pages.ProtocolPage())
	pages.Get("journal", s.pages.JournalPage())

	protocol := s.router.Group("/protocols")
	protocol.Use(logger.New())
	protocol.Use(middlewares.CheckJWT())
	protocol.Get("/:filter/:center", s.handlers.GetProtocols())
	protocol.Patch("/save", s.handlers.SaveProtocol())
	protocol.Post("/add", s.handlers.AddProtocol())
	protocol.Delete("/delete", s.handlers.DeleteProtocol())

	adminPanel := s.router.Group("/admin")
	adminPanel.Get("/panel", s.pages.AdminPage())

	center := s.router.Group("/center")
	center.Use(logger.New())
	center.Use(middlewares.CheckJWT())
	center.Post("/add", s.handlers.AddNewCenter())
	center.Patch("/update", s.handlers.UpdateCenter())
	center.Get("/all", s.handlers.GetCenters())
	center.Get("/name/:id", s.handlers.GetCenterName())
	center.Delete("/delete", s.handlers.DeleteCenter())

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
