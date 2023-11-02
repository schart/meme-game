package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"shared-library/utils"

	"github.com/gofrs/uuid"
)

func PhotoUploadService(w http.ResponseWriter, r *http.Request) error {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("Format of image is corrupted ")
	}
	defer file.Close()

	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")
	if uploadPath == "" {
		return fmt.Errorf("PHOTO_UPLOAD_PATH environment variable is not set")
	}

	filePath = filepath.Join(filePath, uploadPath, handler.Filename)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Error saving the file")
	}
	defer f.Close()

	/*

	  @ Specified max memory for the taken photo and taken photo in body
	  @ Open file with flas of create and write
	  @ Finally, copy file content and sending to rename service
	*/

	_, err = io.Copy(f, file)
	if err != nil {
		return fmt.Errorf("Error copying file content")
	}

	ext := filepath.Ext(handler.Filename)[1:]
	go PhotoRenameservice(handler.Filename, ext)

	return nil
}

func PhotoRenameservice(oldName, ext string) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("Error generating UUID: %v", err)
	}

	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")

	newFilePath := filepath.Join(filePath+uploadPath, id.String()+"."+ext)
	oldFilePath := filepath.Join(filePath+uploadPath, oldName)

	/*

		@ Taken oldname for change new name.
		@ Finally, Renamed of file name.

	*/

	utils.RenamedVariableKeep(id.String() + "." + ext)

	err = os.Rename(oldFilePath, newFilePath)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("Error renaming file: %v", err)
	}

	return nil
}
