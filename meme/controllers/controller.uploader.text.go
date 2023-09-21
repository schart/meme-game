package controllers

import (
	"net/http"
	interceptors "shared-library/interceptors"
	"shared-library/rabbitmq"
	memeTypes "shared-library/types/meme"
	utils "shared-library/utils"
)

func TextUploadController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
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

	// Get parsed body
	formData := utils.ParsedBodyGet(w, r)

	// Check body field with structure
	err := utils.ParameterChecker(formData, memeTypes.TextUpload{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	_interface := memeTypes.TextUpload{
		Text: formData.Get("Text"),
	}

	// Send text for upload to database
	rabbitmq.SendMessage(_interface.Text, "textq")

	// Receive text for upload to database
	rabbitmq.ReceiveText("textq")

	utils.HandleSuccess(w, []string{})
}
