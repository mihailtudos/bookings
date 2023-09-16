package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/mihailtudos/bookings/internal/config"
	"github.com/mihailtudos/bookings/internal/handlers"
	"github.com/mihailtudos/bookings/internal/models"
	"github.com/mihailtudos/bookings/internal/render"
	"log"
	"net/http"
	"time"
)

const POST_NUMBER = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on %s", POST_NUMBER)
	//_ = http.ListenAndServe(POST_NUMBER, nil)

	srv := &http.Server{
		Addr:    POST_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Cannot create the server.", err)
	}
}

func run() error {
	app.InProduction = false
	//allow to store reservations in session
	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create the template cache.", err)
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	return nil
}
