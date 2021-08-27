package main

import (
	"flag"
)

func main()  {
	if len(flag.Args()) == 0 {
		logger.Error.Fatal("Missing music address:(")
	}

	
}