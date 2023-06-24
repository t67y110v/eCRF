package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/t67y110v/web/docs"
	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/handlers"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/store"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger" // swagger handler
)

type server struct {
	router   *fiber.App
	logger   logging.Logger
	pgStore  store.PostgresStore
	mgStore  store.MongoStore
	config   *config.Config
	handlers *handlers.Handlers
}

func newServer(
	pgstore store.PostgresStore,
	mgstore store.MongoStore,
	config *config.Config,
	log logging.Logger,
) *server {

	s := &server{
		router: fiber.New(
			fiber.Config{
				ServerHeader: "eCRF API",
				AppName:      "Api v1.0.1",
			}),
		logger:   log,
		pgStore:  pgstore,
		mgStore:  mgstore,
		config:   config,
		handlers: handlers.NewHandlers(pgstore, mgstore, log),
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
	//user.Use(middlewares.CheckJWT())
	user.Post("/register", s.handlers.UserRegister())
	user.Patch("/update", s.handlers.UserUpdate())
	user.Delete("/delete", s.handlers.UserDelete())
	user.Get("/all", s.handlers.GetUsers())
	//////////////////////////////////////

	protocol := s.router.Group("/protocols")
	protocol.Use(logger.New())
	//	protocol.Use(middlewares.CheckJWT())
	protocol.Get("/:filter/:center", s.handlers.GetProtocols())
	protocol.Patch("/save", s.handlers.SaveProtocol())
	protocol.Post("/add", s.handlers.AddProtocol())
	protocol.Delete("/delete", s.handlers.DeleteProtocol())

	center := s.router.Group("/center")
	center.Use(logger.New())
	//	center.Use(middlewares.CheckJWT())
	center.Post("/add", s.handlers.AddNewCenter())
	center.Patch("/update", s.handlers.UpdateCenter())
	center.Get("/all", s.handlers.GetCenters())
	center.Get("/name/:id", s.handlers.GetCenterName())
	center.Delete("/delete", s.handlers.DeleteCenter())

	subject := s.router.Group("/subject")
	//subject.Use(middlewares.CheckJWT())
	subject.Post("/add", s.handlers.AddSubject())
	subject.Get("/:protocol_id", s.handlers.GetSubjects())
	subject.Get("/:protocol_id/:subject_num", s.handlers.GetSubjectByNumber())
	screening := subject.Group("/screening")
	screening.Patch("/informaionconsent", s.handlers.InformationConsentSubject())
	screening.Patch("/demography", s.handlers.DemographySubject())
	screening.Patch("/anthropometry", s.handlers.AnthropometrySubject())
	screening.Patch("/inclusioncriteria", s.handlers.InclusionCriteriaSubject())
	screening.Patch("/exclusioncriteria", s.handlers.ExclusionСriteriaSubject())
	screening.Patch("/completion", s.handlers.CompletionOfScreening())
	action := subject.Group("/action")
	action.Patch("/updatecolor", s.handlers.UpdateColor())
	action.Patch("/updatecolorwithcomment", s.handlers.UpdateColorWithComment())
	action.Patch("/updatefield", s.handlers.UpdateFieldValue())

	// errors := s.router.Group("/error")
	// errors.Get("/", s.pages.ErrorPage())

}
