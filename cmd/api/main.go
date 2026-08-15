package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amarnathdbg101-coder/olx/internal/config"
	"github.com/amarnathdbg101-coder/olx/internal/handler"
)
func main() {

    config.MustLoad()
     mux := http.NewServeMux()
	 mux.HandleFunc("GET /healthz",handler.Healthz)
    
     srvr := http.Server{
		Addr: ":"+os.Getenv("PORT"),
		Handler: mux,
		ReadTimeout: time.Second*10,
		WriteTimeout: time.Second*30,
		IdleTimeout: time.Second*60,
	 }

	if err := srvr.ListenAndServe(); err != nil {
		log.Fatalf("server failed:%v",err)
	}
}