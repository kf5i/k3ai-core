package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kf5i/k3ai-core/internal/plugins"
)

func main() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.HandleFunc("/plugins", pluginListHandler)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func pluginListHandler(w http.ResponseWriter, r *http.Request) {
	plugins, err := plugins.ContentList(plugins.DefaultPluginURI)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(plugins)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
