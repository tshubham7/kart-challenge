package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tshubham7/kart-challenge/internal/router"
)

func main() {
	r := mux.NewRouter()
	router.AddRoutes(r)

	logrus.Info("starting server :8088...")
	go func() {
		if err := http.ListenAndServe(":8088", r); err != nil {
			logrus.Fatal("failed to start the server, ", err)
		}
	}()
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quitChannel
	logrus.Info("quitting server...")
}
