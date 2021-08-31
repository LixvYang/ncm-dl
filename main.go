package main

import (
	"flag"
	"ncm-dl/common"
	"ncm-dl/logger"
	"ncm-dl/utils"
	"ncm-cl/handler"
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
		logger.Error.Fatal(err)
	}

	if err = req.Do(); err != nil {
		logger.Error.Fatal(err)
	}
}
