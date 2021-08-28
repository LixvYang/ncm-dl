package common

import (
	"time"
)

const (
	NeteaseMusicOrigin  = "https://music.163.com"
	NeteaseMusicReferer = "https://music.163.com"
	NeteaseMusicCookie  = "appver=4.1.3; MUSIC_U=dc5b075d43a3815098b96ce361510c5045495a205ea9ff85e5ea6c502a777644538edaa51b1a56a3e83f5c6c7e24588e4e655b70a75a628ebf122d59fa1ed6a2"
	TencentMusicOrigin  = "https://c.y.qq.com"
	TencentMusicReferer = "https://c.y.qq.com"
	RequestTimeout      = 120 * time.Second
)

type MusicRequest interface {
	Do() error
	Extract() ([]*MP3,error)
}