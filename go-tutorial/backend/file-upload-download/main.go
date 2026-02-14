package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const uploadDir = "./files"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join(uploadDir, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer dst.Close()

	io.Copy(dst, file)
	w.Write([]byte("uploaded"))
}

func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, e.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path)
	filePath := filepath.Join(uploadDir, fileName)

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	http.ServeFile(w, r, filePath)
}

func main() {
	os.MkdirAll(uploadDir, os.ModePerm)

	frontendServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", frontendServer)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", listFilesHandler)
	http.HandleFunc("/download/", downloadHandler)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}