package main

import (
	"flag"
	"ncm-dl/common"
	"ncm-dl/logger"
	"ncm-dl/utils"
	"ncm-dl/handler"
)

func main() {
	if len(flag.Args()) == 0 {
		logger.Error.Fatal("Missing music address:(")
	}

	if err := utils.BuildPathIfNotExist(common.MP3DownloadDir); err != nil {
		logger.Error.Fatalf("Failed to build path: %s: %s", common.MP3DownloadDir, err)
	}

	url := flag.Args()[0]
	req, err := handler.Parse(url)
	if err != nil {
		logger.Error.Fatalf("Failed to build hander Parse:%s",err)
	}

	if err = req.Do(); err != nil {
		logger.Error.Fatalf("Failed to build req Do:%s",err)
	}

	mp3List, err := req.Extract()
	if err != nil {
		logger.Error.Fatalf("Failed to build req Extract:%s",err)
	}

	n := common.MP3ConcurrentDownloadTasksNumber
	for n > common.MaxConcurrentDownloadTasksNumber {
		n = common.MaxConcurrentDownloadTasksNumber
	}
	switch {
	case n > 1:
		handler.ConcurrentDownload(mp3List, n)
	default:
		handler.SingleDownload(mp3List)
	}
}

