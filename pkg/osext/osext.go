package osext

import "os"

func CurrentFolder() string {
	path, _ := os.Getwd()
	return path
}
