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
	checkMethod := utils.HttpMethodSet(http.MethodPost, r)
	if checkMethod != true {
		utils.HandleError(w, http.StatusMethodNotAllowed, "Method error expected method "+http.MethodPost)
		return
	}

	utils.EnvLoader()

	presenceOfKey := interceptors.AccessKeyCheck(w, r)
	if presenceOfKey == false {
		utils.HandleError(w, http.StatusBadRequest, "Error: Required access key, check your 'access key'")
		return
	}

	/*

		@ We checked access key
		@ Before, upload the photo by calling upload service

	*/

	err := memeService.PhotoUploadService(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Error uploading file: "+err.Error())
		return
	}

	/*

		@  Finally, rename file and sent to line up for saving to database

	*/

	renamedFileName := utils.RenamedVariableTurn()
	rabbitmq.SendMessage(renamedFileName, "photoq")
	rabbitmq.ReceivePhotoId("photoq")

	utils.HandleSuccess(w, map[string]interface{}{})
}
