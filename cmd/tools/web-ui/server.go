package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kf5i/k3ai-core/internal/plugins"
)

type httpService struct {
}

func (svc *httpService) Run(stop <-chan struct{}) {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/plugins", pluginListHandler)

	static := r.PathPrefix("/")
	static.Handler(http.StripPrefix("/", http.FileServer(http.Dir("dist"))))

	port := "2200"
	server := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Printf("🚀 Server started at http://localhost:" + port)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_ = server.Shutdown(ctx)
}

func pluginListHandler(w http.ResponseWriter, r *http.Request) {
	plugins, err := plugins.ContentList(plugins.DefaultRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	data, err := json.Marshal(plugins)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func main() {
	stop := make(chan struct{})
	svc := &httpService{}
	go svc.Run(stop)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals
	close(stop)

}
