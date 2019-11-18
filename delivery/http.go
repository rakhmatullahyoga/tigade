package delivery

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func SetupHttpHandler() *mux.Router {
	h := mux.NewRouter()
	h.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	return h
}

func RunHttpServer(h http.Handler, port uint16) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	s := &http.Server{
		Addr: addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h, // Pass our instance of gorilla/mux in.
	}

	go func() {
		log.Printf("The service is running at %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	s.Shutdown(context.Background())
	log.Println("shutting down")
	os.Exit(0)
}
