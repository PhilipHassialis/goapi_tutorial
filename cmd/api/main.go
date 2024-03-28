package main

import (
	"fmt"
	"net/http"

	"github.com/PhilipHassialis/goapi_tutorial/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service...")
	fmt.Println(`
	________  ________          ________  ________  ___     
	|\   ____\|\   __  \        |\   __  \|\   __  \|\  \    
	\ \  \___|\ \  \|\  \       \ \  \|\  \ \  \|\  \ \  \   
	 \ \  \  __\ \  \\\  \       \ \   __  \ \   ____\ \  \  
	  \ \  \|\  \ \  \\\  \       \ \  \ \  \ \  \___|\ \  \ 
	   \ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\
	    \|_______|\|_______|        \|__|\|__|\|__|     \|__|
	`)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
