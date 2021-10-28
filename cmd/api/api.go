package main

import (
	apperror "api/internal/apperrors"
	"api/internal/auth"
	"api/internal/config"
	"api/internal/database"

	"api/internal/logging"
	"api/internal/prom"
	"api/internal/server"
	docs "api/internal/swagger"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

var (
	hostServer = flag.String("host", "0.0.0.0", "Bind addres by API server")
	portServer = flag.String("port", "8000", "Bind port by API server")
	configFile = flag.String("config", "config.yml", "API server configuration file")
)

func main() {
	// Parse CLI arguments
	flag.Parse()

	// Create context
	ctx := context.Background()

	// Create channel for gracefully shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// Read config file
	yfile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		logging.Logger.Fatalf("Read config file error: %s", err.Error())
	}
	conf := config.Config{}
	err = yaml.Unmarshal(yfile, &conf)
	if err != nil {
		logging.Logger.Fatalf("Read YAML file error: %s", err.Error())
	}

	// Create DB connection
	db, err := database.New(&database.Database{
		Host:     conf.Database["host"],
		Port:     conf.Database["port"],
		User:     conf.Database["user"],
		Password: conf.Database["password"],
		DBName:   conf.Database["dbname"],
		SSLMode:  conf.Database["sslmode"],
		Logger:   *logging.Logger,
	})
	if err != nil {
		logging.Logger.Fatalln(err.Error())
	}

	// Create router
	router := mux.NewRouter()
	// Update default errors handlers
	router.MethodNotAllowedHandler = apperror.MethodNotAllowed()
	router.NotFoundHandler = apperror.NotFound()

	// API subroute
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(logging.LoggingMiddleware)

	// Init handlers
	// Prometheus metrics
	promhandler := prom.NewHandler()
	promhandler.Register(router)
	// Docs
	docsHandler := docs.NewHandler()
	docsHandler.Register(router)
	// Auth
	authStorage := auth.NewStorage(db)
	authService := auth.NewAuthService(authStorage)
	vdcHandler := auth.NewHandler(*authService)
	vdcHandler.Register(apiRouter)

	// Create new server and start
	apiServer := server.NewServer(fmt.Sprintf("%s:%s", *hostServer, *portServer), router, logging.Logger)
	go apiServer.Run()

	// waiting for interrupt signal
	<-stop
	// Stop server with timeout
	apiServer.Stop(ctx)
}
