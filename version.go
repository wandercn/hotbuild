package main

import (
	"log"
	"os"
)

var goVersion = "1.16.5"
var Version = ""

func init() {
	if Version == "" {
		v, err := os.ReadFile("./version.txt")
		if err != nil {
			log.Printf("read ./version.txt failed: %v", err)
			return
		}
		Version = string(v)
	}
}
