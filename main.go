package main

import (
	account "account/controllers"
	controllers_game "game/controllers"
	controllers_meme "meme-uploader/controllers"
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
	schemasMeme.GameCreateTables()

	// Host endpoint services

	// r := mux.NewRouter() re-config with mux
	http.HandleFunc("/meme/uploaders/text", controllers_meme.TextUploadController)
	http.HandleFunc("/meme/uploaders/photo", controllers_meme.PhotoUploadController)

	http.HandleFunc("/meme/items/text/{count}", controllers_meme.TextItemsController)
	http.HandleFunc("/meme/items/photo/{count}", controllers_meme.PhotoItemsController)

	http.HandleFunc("/account/register", account.AccountRegisterController)
	http.HandleFunc("/account/login", account.AccountLoginController)
	http.HandleFunc("/account/logout", account.AccountLogoutController)

	http.HandleFunc("/game/create-room", controllers_game.RoomCreateController)
	http.HandleFunc("/game/join-room", controllers_game.JoinRoomController)

	// Start rabbitmq queues
	rabbitmq.QueueRabbitStart()

	// Start server
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), nil)
}
