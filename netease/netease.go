package netease

import (
	"ncm-dl/common"
	"strconv"
	"fmt"
	"regexp"
)

const (
	UrlParameter = "/(song|artist|album|playlist)\\?id=(\\d+)"
)

func Parse(url string) (req common.MusicRequest,err error) {
	re := regexp.MustCompile(UrlParameter)
	matched, ok := re.FindStringSubmatch(url), re.MatchString(url)
	if !ok || len(matched) < 3 {
		err = fmt.Errorf("could not parse the url: %s", url)
	}
	
	id, err := strconv.Atoi(matched[2])
	if err != nil {
		return
	}

	switch matched[1] {
	case "song":
		req = NewSongRequest(id)
	}

	return
}