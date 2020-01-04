package file

import (
	"net/http"
	"os"
)

type (
	URL      string
	Uploader interface {
		UploadFile(f *os.File) (string, error)
	}
)

func GetFileContentType(f *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := f.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
