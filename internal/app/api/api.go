package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// базовый API сервер
type API struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// API конструктор
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// старт сервера, конфигурация логгера, роутера и БД...
func (api *API) Start() error {
	// конфигурация логгера
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// подтверждение, что логгер сконфигерирован
	api.logger.Info("starting api server at port:", api.config.BindAddr)

	// конфигурация маршрутизатора
	api.configureRouterField()

	return http.ListenAndServe(api.config.BindAddr, api.router)
}
