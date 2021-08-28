package common

import (
)

type Tag struct {
	Title      string
	Artist     string
	Album      string
	Year       string
	Track      string
	CoverImage string
}

type MP3 struct {
	FileName string
	SavePath    string
	Playable    bool
	DownloadUrl string
	Tag         Tag
	Origin      int
}