package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// конфигирация API instance (поле logger)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// конфигурация маршрутизатора (поле router API)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is REST API"))
	})
}
