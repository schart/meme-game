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

// File photo upload to disk
func PhotoUploadService(w http.ResponseWriter, r *http.Request) error {
	// Maximum image size length can be 10
	r.ParseMultipartForm(10 << 20)

	// Get file
	file, handler, err := r.FormFile("file")
	if err != nil {
		return fmt.Errorf("Format of image is corrupted ")
	}
	defer file.Close()

	// Get path of ll upload disk
	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")
	if uploadPath == "" {
		return fmt.Errorf("PHOTO_UPLOAD_PATH environment variable is not set")
	}

	// Join path for processing
	filePath = filepath.Join(filePath, uploadPath, handler.Filename)

	// Open the file with flags
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("Error saving the file")
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return fmt.Errorf("Error copying file content")
	}

	// Rename photo
	ext := filepath.Ext(handler.Filename)[1:]
	go PhotoRenameservice(handler.Filename, ext)

	return nil
}

// Rename photo with id
func PhotoRenameservice(oldName, ext string) error {
	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("Error generating UUID: %v", err)
	}

	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")

	newFilePath := filepath.Join(filePath+uploadPath, id.String()+"."+ext)
	oldFilePath := filepath.Join(filePath+uploadPath, oldName)

	utils.RenamedVariableKeep(id.String() + "." + ext)

	err = os.Rename(oldFilePath, newFilePath)
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("Error renaming file: %v", err)
	}

	return nil
}
