package reloader

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func Reloader() {
	// Create a watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error wather: ", err)
	}

	// Do this process lastly
	defer watcher.Close()

	done := make(chan bool)
	fmt.Println("value of done: ", &done)
	go func() {
		for {
			// evulation of events
			select {
			// watch events
			case event := <-watcher.Events:
				fmt.Printf("Event: %v", event.Op)

			case err := <-watcher.Errors:
				fmt.Printf("Event: %v", err)

			}
		}
	}()

	path, err := os.Getwd()
	if err != nil {
		fmt.Errorf("Error Path not found: ", err)
	}

	path = path + "\\"

	if err := watcher.Add(path); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
	fmt.Println("value of done lastly: ", &done)
}
