package main

import (
	"log"
	"runtime/debug"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal("no build info")
	}

	var rev, time, dirty string
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			rev = setting.Value
		case "vcs.time":
			time = setting.Value
		case "vcs.modified":
			dirty = setting.Value
		}
	}
	log.Println("Revision:", rev)
	log.Println("Time:", time)
	log.Println("Dirty:", dirty)
}
