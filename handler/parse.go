package handler

import (
	"fmt"
	"ncm-dl/common"
	"ncm-dl/logger"
	"ncm-dl/netease"
	"regexp"
)

const (
	UrlPattern = "music.163|y.qq.com"
)

func Parse(url string) (req common.MusicRequest,err error) {
	re := regexp.MustCompile(UrlPattern)
	matched,ok := re.FindString(url),re.MatchString(url)
	if !ok {
		err = fmt.Errorf("could not parse the url: %s", url)
		return
	}

	switch matched {
	case "music.163.com":
		req, err = netease.Parse(url)
		if err != nil {
			logger.Error.Fatal("could not matched url")
		}
	}

	return req, err
}
