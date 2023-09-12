package main

// my-monorepo/main.go

import (
	"fmt"
	uploaders "meme-uploader/controllers"
	"net/http"
	"os"
	schemasMeme "shared-library/database/schemas"

	"shared-library/rabbitmq"
	"shared-library/utils"
)

func main() {
	utils.EnvLoader() // Load env file information

	schemasMeme.MemeCreateT()
	// dirproccessors.JsonReader()

	// Host endpoint services
	http.HandleFunc("/uploaders/upload-text", uploaders.TextUploadController)
	http.HandleFunc("/uploaders/upload-photo", uploaders.PhotoUploadController)

	// Watch directory
	// go listeners.PhotoDirectoryListener()
	// go listeners.TextDirectoryListener()

	// Start rabbitmq queues
	rabbitmq.QueueRabbitStart()

	// Restart server on the change
	// go reloader.Reloader()

	// Start server
	fmt.Printf("Server started on the %v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
