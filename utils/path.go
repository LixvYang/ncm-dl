package utils

import (
	"os"
)

//
func ExistsPath(path string) (bool,error) {
	_, err := os.Stat(path)
	if err != nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return true,nil 
}

//check if path exist
func BuildPathIfNotExist(path string) error {
	ok,err := ExistsPath(path)
	if !ok {
		return os.MkdirAll(path,0644)
	} else { 
		return err
	}
}
