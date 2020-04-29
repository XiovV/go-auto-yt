package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/xiovv/go-auto-yt/models"
	"github.com/xiovv/go-auto-yt/orm"
	"net/http"
)

func HandleChannelDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/channels.html")
}

func AddChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var targetPreferences models.TargetDownloadPreferences
	_ = json.NewDecoder(r.Body).Decode(&targetPreferences)

	orm.Insert(&targetPreferences)
}

func GetChannelMetadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var target models.TargetDownloadPreferences
	_ = json.NewDecoder(r.Body).Decode(&target)

	fmt.Println(target)
	metadata, _ := target.GetMetadata()
	fmt.Println(metadata.PlaylistUploader)
	fmt.Println(metadata.Title)
	fmt.Println(metadata.WebpageURL)

	response := make(map[string]string)
	response["channelName"] = metadata.PlaylistUploader
	response["latestVideo"] = metadata.Title
	response["latestVideoURL"] = metadata.WebpageURL

	json.NewEncoder(w).Encode(&response)
}