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

	utils.EnvLoader()

	presenceOfKey := interceptors.AccessKeyCheck(w, r)
	if presenceOfKey == false {
		utils.HandleError(w, http.StatusBadRequest, "Error: Required access key, check your 'access key'")
		return
	}

	formData := utils.ParsedBodyGet(w, r)
	err := utils.ParameterChecker(formData, memeTypes.TextUpload{})
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	/*

		@ We checked access key, body parameter
		@ Finally, sending line up for saving database

	*/

	_interface := memeTypes.TextUpload{
		Text: formData.Get("Text"),
	}

	rabbitmq.SendMessage(_interface.Text, "textq")
	rabbitmq.ReceiveText("textq")

	utils.HandleSuccess(w, map[string]interface{}{})
}
