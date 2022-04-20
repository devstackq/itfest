package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"bimbo/internal/config"
	"bimbo/internal/handler"

	"bimbo/internal/repository/psql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type App struct {
	// grpc        grpc.Server
	// authUseCase auth.UseCase
	cfg    *config.Config
	db     *sql.DB
	router *gin.Engine
	Logger *logrus.Logger
}

// interface {Signup, Signin}; stuct Grpc - own realize; struct http - own realize, grpcServer
// singletone - prepare app, connect layers with interface; init app

func NewApp(cfg *config.Config) (*App, error) {
	if cfg == nil {
		return nil, fmt.Errorf("error app struct")
	}
	return &App{cfg: cfg}, nil
}

// @title        Bazar service
// @version      1.0

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @BasePath  /
// @schemes   http

func (a *App) Initialize() {
	gin.SetMode(a.cfg.App.Mode)

	a.router = gin.New()
	a.Logger = logrus.New()

	a.router.Use(gin.Recovery())
	a.router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		MaxAge:        30 * time.Second,
		AllowWildcard: true,
	}))

	// mongoObject := db.NewDbObject("mongodb", viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.host"), viper.GetString("mongo.port"), viper.GetString("mongo.dbName"), viper.GetString("mongo.user_collection"))
	// repo := mongoRepo.NewUserRepository(db.(*mongo.Database), viper.GetString("mongo.user_collection"))

	db, err := psql.InitDb(a.cfg)
	if err != nil {
		log.Println(err)
		// return
	}
	a.db = db
	a.Logger.Info("intialize postgres...")
	a.setComponents()
}

func (a *App) Run(ctx context.Context) {
	srv := http.Server{
		Addr:           a.cfg.App.Port,
		Handler:        a.router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    a.cfg.App.ReadTimeout,
		WriteTimeout:   a.cfg.App.WriteTimeout,
	}
	go func() {
		a.Logger.Info("starting web server on port: ", a.cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Fatal(err.Error())
		}
	}()

	<-ctx.Done()

	a.Logger.Info("shutting down web server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.Logger.Fatal("application forced to shutdown: ", err.Error())
	}
	a.Logger.Info("application exiting")
}

func (a *App) setComponents() {
	apiVersion := a.router.Group("/v1")
	apiVersion.Static("/images/", "./images")
	handler.SetEndpoints(a.cfg, a.db, a.Logger, apiVersion)
}
