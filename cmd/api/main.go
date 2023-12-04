package main

import (
	"fmt"
	"net/http"

	"github.com/gaba-bouliva/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	

	fmt.Println("Starting GO API service...")
	fmt.Println(`
	   ____            _    ____ ___ 
  / ___| ___      / \  |  _ \_ _|
 | |  _ / _ \    / _ \ | |_) | | 
 | |_| | (_) |  / ___ \|  __/| | 
  \____|\___/  /_/   \_\_|  |___|

	`)
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}