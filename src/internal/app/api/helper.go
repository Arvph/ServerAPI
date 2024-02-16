package api

import (
	"net/http"

	"github.com/arvph/ServerAPI/storage"
	"github.com/sirupsen/logrus"
)

// API congig instance (logger field)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Router configs (router field)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is REST API"))
	})
}

// Storage configs
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
