package logger

import (
	"os"
	"log"
)

var (
	Debug = log.New(os.Stderr,"[Error]",log.LstdFlags)
	Info    = log.New(os.Stdout, "[Info] ", log.LstdFlags)
	Warning = log.New(os.Stdout, "[Warning] ", log.LstdFlags)
	Error   = log.New(os.Stderr, "[Error] ", log.LstdFlags)
)