package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

func (target DownloadTarget) GetMetadata() (TargetMetadata, error) {
	cmd := exec.Command("youtube-dl", "-j", "--playlist-end", "1", target.URL)
	log.Info("executing youtube-dl command: ", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("From GetMetadata(): ", err)
		return TargetMetadata{}, fmt.Errorf("From c.GetMetadata(): %v", string(out))
	}
	metaData := &TargetMetadata{}
	if err = json.Unmarshal(out, metaData); err != nil {
		log.Error("From GetMetadata(): ", err)
		return TargetMetadata{}, fmt.Errorf("From c.GetMetadata(): %v", string(out))
	}

	return *metaData, nil
}

func (target DownloadTarget) GetLatestVideo() (string, error) {
	log.Info("fetching latest upload")
	cmd := exec.Command("youtube-dl", "-j", "--playlist-end", "1", target.URL)
	log.Info("executing youtube-dl command: ", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error(string(out))
		log.Errorf("c.GetLatestVideo: %s | %s", err, string(out))
	}
	metaData := &TargetMetadata{}
	if err = json.Unmarshal(out, metaData); err != nil {
		log.Error("c.GetLatestVideo: ", err)
		return "", fmt.Errorf("c.GetLatestVideo: %s", err)
	}
	log.Info("successfully fetched latest video ")
	return metaData.ID, nil
}

func (target DownloadTarget) Download(downloadQuality, fileExtension string, downloadEntire, isCustom bool) error {
	log.Info("DOWNLOAD: ", target)
	var ytdlCommand YTDLCommand
	if target.DownloadMode == "Audio Only" {
		log.Info("downloading audio only")
		switch downloadQuality {
		case "best":
			downloadQuality = "0"
		case "medium":
			downloadQuality = "5"
		case "worst":
			downloadQuality = "9"
		}
		log.Info("download quality set to: ", downloadQuality)
	}
	ytdlCommand = YTDLCommand{
		Binary: "youtube-dl",
		Target: target.URL,
	}
	switch target.DownloadMode {
	case "Audio Only":
		ytdlCommand.FileType = "bestaudio"
		ytdlCommand.Output = filepath.Join(dlRoot, target.DownloadPath)
	case "Video And Audio":
		ytdlCommand.FileType = "bestvideo[height<=" + downloadQuality + "]" + "+ bestaudio/best[height<=" + downloadQuality + "]"
		ytdlCommand.Output = filepath.Join(dlRoot, target.DownloadPath)
	}
	switch downloadEntire {
	case true:
		ytdlCommand.FirstFlag = "--ignore-errors"
	case false:
		ytdlCommand.FirstFlag = "--playlist-end"
		ytdlCommand.FirstFlagArg = "1"
	}
	switch isCustom {
	case true:
		ytdlCommand.Custom = target.DownloadPath + " " + target.URL
	}

	fmt.Println("CUSTOM COMMAND: ", target.DownloadPath)

	err := DownloadVideo(ytdlCommand)
	if err != nil {
		return fmt.Errorf("Download: %v", err)
	}
	return nil
}

func DownloadVideo(command YTDLCommand) error {
	log.Info("downloading video")
	var cmd *exec.Cmd
	log.Info(command)
	if command.Custom != "" {
		fmt.Println("RUN CUSTOM COMMAND")
		ytdlCmd := strings.Split(command.Custom, " ")
		cmd = exec.Command(ytdlCmd[0], ytdlCmd[1:]...)
	} else if command.FirstFlag == "" {
		fmt.Println("RUN BASIC CMD 1")
		cmd = exec.Command(command.Binary, "-f", command.FileType, "-o", command.Output, command.Target)
	} else if command.FirstFlag != "" {
		fmt.Println("RUN BASIC CMD 2")
		cmd = exec.Command(command.Binary, command.FirstFlag, command.FirstFlagArg, "-f", command.FileType, "-o", command.Output, command.Target)
	}
	log.Info("executing youtube-dl command: ", cmd.String())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("There was an error downloading the video: ", string(out))
		return fmt.Errorf("DownloadVideo: %s", string(out))
	}
	log.Info(string(out))
	return nil
}
