package api

import (
	"net/http"

	"github.com/arvph/ServerAPI/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Base API server instance description
type API struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server, configure loggers, router, database connection and etc...
func (api *API) Start() error {
	// Trying to configure Logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}

	api.logger.Info("starting api server at port:", api.config.BindAddr)

	// router configs
	api.configureRouterField()
	// storage configs
	if err := api.configureStorageField(); err != nil {
		return nil
	}
	// http server starts
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
