// Package netease provides
package netease

import (
	"encoding/json"
	"fmt"
	"ncm-dl/common"
	"ncm-dl/logger"
	"net/http"
	"net/url"
	"strings"
)

const (
	WeAPI       = "https://music.163.com/weapi"
	SongUrlAPI  = WeAPI + "/song/enhance/player/url"
	SongAPI     = WeAPI + "/v3/song/detail"
)

type SongUrlParams struct {
	Ids string `json:"ids"`
	Br  int    `json:"br"`
}

type SongUrlResponse struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data []SongUrl `json:"data"`
}

type SongUrlRequest struct {
	Params   SongUrlParams
	Response SongUrlResponse
}


type SongParams struct {
	C string `json:"c"`
}

type SongResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Songs []Song `json:"songs"`
}

type SongRequest struct {
	Params   SongParams
	Response SongResponse
}

func NewSongUrlRequest(ids ...int) *SongUrlRequest {
	br := common.MP3DownloadBr
	switch br {
	case 128, 192, 320:
		br *= 1000
	default:
		br = 999 * 1000
	}
	enc, _ := json.Marshal(ids)
	return &SongUrlRequest{Params: SongUrlParams{Ids: string(enc), Br: br}}
}


func NewSongRequest(ids ...int) *SongRequest {
	c := make([]map[string]int,0,len(ids))
	for _,id := range ids {
		c = append(c,map[string]int{"id":id})
	}

	enc,err := json.Marshal(c)
	if err != nil {
		logger.Error.Fatalf("Json marshal error:%s",err)
	}
	return &SongRequest{Params: SongParams{C: string(enc)}}
}

func (s *SongUrlRequest) Do() error {
	enc, err := json.Marshal(s.Params)
	if err != nil {
		logger.Error.Fatalf("Json marshal error:%s",err)
	}
	params, encSecKey, err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params", params)
	form.Set("encSecKey", encSecKey)
	resp, err := common.Request("POST", SongUrlAPI, nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		logger.Error.Fatalf("common Request failed: %s",err)
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s.Response); err != nil {
		return err
	}
	if s.Response.Code != http.StatusOK {
		return fmt.Errorf("%s %s error: %d %s", resp.Request.Method, resp.Request.URL.String(), s.Response.Code, s.Response.Msg)
	}

	return nil
}

func (s *SongRequest) Do() error {
	enc, err := json.Marshal(s.Params)
	if err != nil {
		logger.Error.Fatalf("Json could not marshal: %s",err)
	}
	params,encSecKey,err := Encrypt(enc)
	if err != nil {
		return err
	}

	form := url.Values{}
	form.Set("params",params)
	form.Set("encSecKey",encSecKey)
	resp, err := common.Request("POST", SongAPI, nil, strings.NewReader(form.Encode()), common.NeteaseMusic)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&s.Response);err != nil {
		return err
	}

	if s.Response.Code != http.StatusOK{
		return fmt.Errorf("%s %s error : %d %s",resp.Request.Method, resp.Request.URL.String(), s.Response.Code, s.Response.Msg)
	}

	return nil
}

func (s *SongRequest) Extract() ([]*common.MP3,error) {
	return ExtractMP3List(s.Response.Songs,".")
}

