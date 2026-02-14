package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// Storage folder path
const uploadDir = "./files"


// Handle file upload request from frontend
func uploadHandler(w http.ResponseWriter, r *http.Request) {

	// Extract file from the request
	file, header, err := r.FormFile("file") // File stream, metadata stored
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()

	// Create destination file: ./files/<fileName>
	dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer dst.Close()

	// Copy file data using chunk-by-chunk transfer
	io.Copy(dst, file)

	// Send response
	w.Write([]byte("uploaded"))
}


// Handle file list request from frontend
func listFilesHandler(w http.ResponseWriter, r *http.Request) {

	// Get all the files inside 'files/'
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Filter directories. Only actual files are included
	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}

	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Convert file list to JSON and send response
	json.NewEncoder(w).Encode(files)
}


// Handle file download request from frontend
func downloadHandler(w http.ResponseWriter, r *http.Request) {

	// Extract file name from the request route
	fileName := filepath.Base(r.URL.Path)

	// Creates full path for the file- ./files/<fileName>
	filePath := filepath.Join(uploadDir, fileName)

	// Set response header- Tell browser to download the file instead of displaying
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	
	// Sends the file at filePath to client. Streams file contents
	http.ServeFile(w, r, filePath)
}


func main() {

	// Create the storage folder if it doesn't exist
	os.MkdirAll(uploadDir, os.ModePerm)

	// Handler for serving static frontend files
	frontendServer := http.FileServer(http.Dir("./static"))

	// Mapping "/" route to frontend handler
	http.Handle("/", frontendServer)

	// Mapping upload, list files, download endpoints to their handlers
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", listFilesHandler)
	http.HandleFunc("/download/", downloadHandler)

	// Starting server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}