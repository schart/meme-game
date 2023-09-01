package controllers

import (
	memeService "meme-uploader/services"
	"net/http"
	"path/filepath"
	"shared-library/rabbitmq"
	utils "shared-library/utils"
)

func TextUploadC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, meme upload of text"))
}

func PhotoUploadC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	utils.EnvLoader()

	// Save file to disk
	fileName, err := memeService.PhotoUploadS(w, r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Error uploading file: "+err.Error())
		return
	}

	// Rename file
	ext := filepath.Ext(fileName)[1:] // Take extension of file
	fileName, err = memeService.PhotoRename(fileName, ext)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Save file to database
	rabbitmq.SendMessage(fileName, "photoq")

	utils.HandleSuccess(w)
}
