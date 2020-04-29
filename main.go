package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	h "github.com/xiovv/go-auto-yt/handlers"
	"os"
)

func createDir(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		f, _ := os.Create(dir)
		defer f.Close()
		f.WriteString("[]")
		f.Sync()
	}
}

func initDatabase() {
	log.Info("Initiating db")
	databases := [3]string{"./config/channels.json", "./config/playlists.json", "./config/videos.json"}
	for _, database := range databases {
		_, err := os.Stat(database)
		if os.IsNotExist(err) {
			f, _ := os.Create(database)
			defer f.Close()
			f.WriteString("[]")
			f.Sync()
		}
	}
}

func init() {
	initDatabase()
}

func main() {
	r := mux.NewRouter()
	//r.Use(h.ContentTypeMiddleware)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	r.HandleFunc("/", h.HandleChannelDashboard).Methods("GET")
	r.HandleFunc("/api/channel/add", h.AddChannel).Methods("POST")
	r.HandleFunc("/api/channel/metadata", h.GetChannelMetadata).Methods("POST")
	_ = http.ListenAndServe(":8080", r)
}
