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

	// Host endpoint services
	http.HandleFunc("/uploaders/upload-text", uploaders.TextUploadC)
	http.HandleFunc("/uploaders/upload-photo", uploaders.PhotoUploadC)

	fmt.Printf("Server started on the %v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	fmt.Println()

	// Watch directory
	// go listeners.PhotoDirectoryListener()
	// go listeners.TextDirectoryListener()

	// Start rabbitmq queues
	rabbitmq.QueueRabbitStart()
	go rabbitmq.ReceivePhotoId("photoq")

	// Start server
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
