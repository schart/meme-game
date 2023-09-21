package main

import (
	account "account/controllers"
	memes "meme-uploader/controllers"
	"net/http"
	"os"
	schemasMeme "shared-library/database/schemas"
	"shared-library/rabbitmq"
	"shared-library/utils"
)

func main() {
	utils.EnvLoader() // Load env file information

	schemasMeme.MemeCreateTables()
	schemasMeme.AccountCreateTables()

	// Host endpoint services

	http.HandleFunc("/meme/uploaders/text", memes.TextUploadController)
	http.HandleFunc("/meme/uploaders/photo", memes.PhotoUploadController)

	http.HandleFunc("/meme/items/text", memes.TextItemsController)
	http.HandleFunc("/meme/items/photo", memes.PhotoItemsController)

	http.HandleFunc("/account/register", account.AccountRegister)
	http.HandleFunc("/account/login", account.AccountLogin)

	// Start rabbitmq queues
	rabbitmq.QueueRabbitStart()

	// Start server
	// fmt.Printf("Server started on the %v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
