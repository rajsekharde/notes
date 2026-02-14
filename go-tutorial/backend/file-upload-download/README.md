# File upload/download using Go backend

Simple full stack web app built entirely using Go's standard net/http package. It allows users to:
- Upload files from a browser
- Store files locally inside /files directory
- Display a list of uploaded files on a web page
- Download files through the browser

## Architecture

Browser(Frontend) -> Go HTTP server -> 'files/' (storage)

API endpoints:
- GET "/" : Homepage. Servers frontend HTML page
- POST "/upload" : Upload file to server
- GET "/files" : List files in JSON format
- GET "/download/[filename]" : Download file from server

## Flow of events

When application starts:
- Server ensures files/ directory exists
- HTTP handlers are registered for all routes
- HTTP server starts listening on port 8080

### Upload files:

Frontend:
- A file is selected using the input field and "Upload" button is clicked
- JavaScript creates a FormData object
- File is appended under the key "file"
- A POST request is sent to /upload
- File list is refreshed automatically after upload

Backend:
- Server receives POST request
- Extracts uploaded file using r.FormFile("file")
- Provided file stream and metadata
- Creates destination file with path- ./files/[filename]
- File data is copied using streamed transfer
- Success response is returned to the frontend

### Listing upload files:

/files endpoint:
- Reads the files/ directory
- Filters out sub-directories
- Creates a list of filenames
- Returns JSON response

Frontend: Uses JSON response to display download links

### Download files:

/download/[filename] handler:
- Extracts file name and creates download path
- Forces browser to download instead of preview by setting Content-Disposition: attachment in response header
- File is served safely using file streaming

## Concurrency handling

Go's HTTP server automatically:
- Spawns a goroutine per request
- Allows multiple uploads/downloads simultaenously
- Requires no manual threading logic