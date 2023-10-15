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

	"github.com/gorilla/mux"
)

func main() {
	utils.EnvLoader() // Load env file information

	schemasMeme.MemeCreateTables()
	schemasMeme.AccountCreateTables()
	schemasMeme.GameCreateTables()

	// Host endpoint services

	// Create a new mux router.
	r := mux.NewRouter()

	// Register each route with the mux router.
	r.HandleFunc("/meme/uploaders/text", controllers_meme.TextUploadController)
	r.HandleFunc("/meme/uploaders/photo", controllers_meme.PhotoUploadController)

	r.HandleFunc("/meme/items/text/{count}", controllers_meme.TextItemsController)
	r.HandleFunc("/meme/items/photo/{count}", controllers_meme.PhotoItemsController)

	r.HandleFunc("/account/register", account.AccountRegisterController)
	r.HandleFunc("/account/login", account.AccountLoginController)
	r.HandleFunc("/account/logout", account.AccountLogoutController)
	r.HandleFunc("/account/items/room", account.GetRoomAccountController)

	r.HandleFunc("/game/create-room", controllers_game.RoomCreateController)
	r.HandleFunc("/game/join-room", controllers_game.JoinRoomController)
	r.HandleFunc("/game/items/rooms", controllers_game.GetAllRoomsController)
	r.HandleFunc("/game/start-game/{room_link}", controllers_game.StartGameController)

	// Start rabbitmq queues
	rabbitmq.QueueRabbitStart()

	// Start server
	http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), r)
}
