package main

import (
	"flag"
	"ncm-dl/logger"
	"ncm-dl/utils"
	"ncm-dl/common"
)

func main() {
	if len(flag.Args()) == 0 {
		logger.Error.Fatal("Missing music address:(")
	}

	if err := utils.BuildPathIfNotExist(common.MP3DownloadDir); err != nil {
		logger.Error.Fatalf("Failed to build path: %s: %s", common.MP3DownloadDir, err)
	}


}
