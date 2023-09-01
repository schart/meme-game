package directorylisteners

import (
	"log"
	"os"
	utils "shared-library/utils"

	"github.com/fsnotify/fsnotify"
)

func PhotoDirectoryListener() {
	utils.EnvLoader()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	filePath, _ := os.Getwd()
	photoUploadPath := os.Getenv("PHOTO_UPLOAD_PATH")

	// Choose the directory and add the listener to there
	err = watcher.Add(filePath + photoUploadPath)
	if err != nil {
		log.Fatal(err)
	}

	// Listen events
	for {
		select {
		case event := <-watcher.Events:
			log.Println("Event:", event.Op)

		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}
