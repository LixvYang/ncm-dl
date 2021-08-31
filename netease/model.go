package netease

import (
)

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Album struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PicUrl      string `json:"picUrl"`
	PublishTime int64  `json:"publishTime"`
}

type SongUrl struct {
	Id   int    `json:"id"`
	Code int    `json:"code"`
	Url  string `json:"url"`
}

type Song struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Artist      []Artist `json:"ar"`
	Album       Album    `json:"al"`
	Position    int      `json:"no"`
	PublishTime int64    `json:"publishTime"`
}

type TrackId struct {
	Id int `json:"id"`
}

type Playlist struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	TrackIds []TrackId `json:"trackIds"`
}