package logger

import (
	"os"
	"log"
)

var (
	Error = log.New(os.Stderr,"[Error]",log.LstdFlags)
)