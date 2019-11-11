package main

import (
	"fmt"
)

// Download downloads a the latest video based on downloadMode
func (c ChannelTest) Download(downloadMode string) error {
	channelName := c.Name
	channelType := c.Type
	if downloadMode == "Video And Audio" {
		// Download .mp4 with audio and video in one file
		return downloadVideoAndAudio(channelName, channelType)
		// Extract audio from the .mp4 file and remove the .mp4
	} else if downloadMode == "Audio Only" {
		return downloadAudioOnly(channelName, channelType)
	}
	return fmt.Errorf("Something went seriously wrong")
}

func downloadVideoAndAudio(channelName, channelType string) error {
	video := GetLatestVideo(channelName, channelType)
	err := video.DownloadYTDL()
	if err != nil {
		return err
	}

	return UpdateLatestDownloaded(channelName, video.VideoID)
}

func downloadAudioOnly(channelName, channelType string) error {
	video := GetLatestVideo(channelName, channelType)
	err := video.DownloadAudioYTDL()
	if err != nil {
		return err
	}

	return UpdateLatestDownloaded(channelName, video.VideoID)
}
