package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/t67y110v/web/internal/app/config"
	"github.com/t67y110v/web/internal/app/logging"
	centerModlel "github.com/t67y110v/web/internal/app/model/center"
	protocolModel "github.com/t67y110v/web/internal/app/model/protocol"
	userModel "github.com/t67y110v/web/internal/app/model/user"

	"github.com/t67y110v/web/internal/app/store/nosqlstore"
	store "github.com/t67y110v/web/internal/app/store/sqlstore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start(config *config.Config) error {

	ctx := context.TODO()

	pgdb, err := newPostgresDB(config)
	if err != nil {
		return err
	}

	mgdb, err := newMongoDB(config, ctx)
	if err != nil {
		return err
	}

	defer mgdb.Disconnect(ctx)

	pgStore := store.NewPostgresDB(pgdb)

	mgStore := nosqlstore.NewMongoDB(mgdb)

	logger := logging.GetLogger()

	server := newServer(pgStore, mgStore, config, logger)

	//StartServerWithGracefulShutdown(server, config.BindAddr)
	return server.router.Listen(config.Port)
}

func newPostgresDB(c *config.Config) (*gorm.DB, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.PG_DBUser, c.PG_DBPass, c.PG_DBHost, c.PG_DBPort, c.PG_DBName)
	log.Printf("Database initialization: %s\n", url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	//db, err := sql.Open("postgres", postgresDatabaseURL)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&protocolModel.Protocol{})
	db.AutoMigrate(&centerModlel.Center{})
	return db, nil
}

func newMongoDB(c *config.Config, ctx context.Context) (*mongo.Client, error) {
	URI := fmt.Sprintf("%s://%s:%s/", c.MG_DBUser, c.MG_DBHost, c.MG_DBPort)
	log.Printf("Database initialization: %s\n", URI)
	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil

}

func StartServerWithGracefulShutdown(s *server, addr string) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := s.router.Shutdown(); err != nil {
			s.logger.Warningf("Server is not shutting down! reason: %v", err)
		}
		close(idleConnsClosed)
	}()
	if err := s.router.Listen(addr); err != nil {
		s.logger.Warningf("Server is not running! reason: %v", err)
	}
	<-idleConnsClosed

}
