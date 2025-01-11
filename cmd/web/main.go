package main

import (
	"fmt"
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"

	// "hellopk/pkg/routes"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false //localhost only

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //localhost only

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cnnot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	fmt.Printf("starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: rooutes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
