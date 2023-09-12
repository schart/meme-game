package controllers

import (
	memeService "meme-uploader/services"
	"net/http"
	interceptors "shared-library/interceptors"
	"shared-library/rabbitmq"
	utils "shared-library/utils"
)

func PhotoUploadController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")
	checkStatus := utils.HttpMethodSet(http.MethodPost, r)
	if checkStatus != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}
	// Load enviroments
	utils.EnvLoader()

	// Check access key
	presenceOfKey := interceptors.AccessKeyCheck(w, r)
	if presenceOfKey == false {
		utils.HandleError(w, http.StatusBadRequest, "Error: Required access key, check your 'access key'")
		return
	}

	// Save file to disk
	err := memeService.PhotoUploadService(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Error uploading file: "+err.Error())
		return
	}

	// Get the renamed name
	renamedFileName := utils.RenamedVariableTurn()

	// Send renamed file name for save to database
	rabbitmq.SendMessage(renamedFileName, "photoq")

	// Receive and upload to db name of photo
	rabbitmq.ReceivePhotoId("photoq")

	utils.HandleSuccess(w)
}
