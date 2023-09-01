package directorylisteners

import (
	"log"
	"os"
	utils "shared-library/utils"

	"github.com/fsnotify/fsnotify"
)

func TextDirectoryListener() {
	utils.EnvLoader()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	filePath, _ := os.Getwd()
	textUploadPath := os.Getenv("TEXT_UPLOAD_PATH")

	// Choose the directory and add the listener to there
	err = watcher.Add(filePath + textUploadPath)
	if err != nil {
		log.Fatal(err)
	}

	// Listen events
	for {
		select {
		case event := <-watcher.Events:
			log.Println("Event:", event)

		case err := <-watcher.Errors:
			log.Println("Error:", err)
		}
	}
}
