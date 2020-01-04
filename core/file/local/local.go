package local

import (
	"os"
)

type (
	Uploader struct {
	}
)

func (u *Uploader) UploadFile(f *os.File) (string, error) {
	return "", nil
}

func NewUploader() *Uploader {
	return &Uploader{}
}
