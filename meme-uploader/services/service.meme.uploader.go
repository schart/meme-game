package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofrs/uuid"
)

// File upload to disk and database
func PhotoUploadS(w http.ResponseWriter, r *http.Request) (string, error) {
	// Maximum image size length can be 10
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		return "", fmt.Errorf("Format of image is corrupted")
	}
	defer file.Close()

	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")
	if uploadPath == "" {
		return "", fmt.Errorf("PHOTO_UPLOAD_PATH environment variable is not set")
	}

	filePath = filepath.Join(filePath, uploadPath, handler.Filename)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", fmt.Errorf("Error saving the file")
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", fmt.Errorf("Error copying file content")
	}

	return handler.Filename, nil
}

// Rename photo with id
func PhotoRename(oldName, ext string) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("Error generating UUID: %v", err)
	}

	filePath, _ := os.Getwd()
	uploadPath := os.Getenv("PHOTO_UPLOAD_PATH")

	newFilePath := filepath.Join(filePath+uploadPath, id.String()+"."+ext)
	oldFilePath := filepath.Join(filePath+uploadPath, oldName)

	err = os.Rename(oldFilePath, newFilePath)
	if err != nil {
		return "", fmt.Errorf("Error renaming file: %v", err)
	}

	fullName := id.String() + "." + ext
	return fullName, nil
}
