package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ishvaram/betal-kamailio/driver"
	subs "github.com/ishvaram/betal-kamailio/handler/http"
)

func main() {
	dbHost := "127.0.0.1"
	dbUserName := "root"
	dbPass := "exotel"
	dbPort := "3306"
	dbName := "kamailio"

	connection, err := driver.ConnectSQL(dbHost, dbPort, dbUserName, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	sHandler := subs.NewSubscriberHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/subscriber", subscriberRouter(sHandler))
	})

	fmt.Println("Server listen at :8026")
	http.ListenAndServe(":8026", r)
}

// subscriberRouter - A completely separate router for subscriber routes
func subscriberRouter(subsHandler *subs.Subscriber) http.Handler {
	r := chi.NewRouter()
	r.Get("/", subsHandler.Fetch)
	r.Get("/{id:[0-9]+}", subsHandler.GetByID)
	r.Post("/", subsHandler.Create)
	r.Put("/{id:[0-9]+}", subsHandler.Update)
	r.Delete("/{id:[0-9]+}", subsHandler.Delete)
	return r
}
