package local

import (
	"context"
	"os"
)

type (
	Uploader struct {
	}
)

func (u *Uploader) UploadFile(ctx context.Context, f *os.File) (string, error) {
	return "", nil
}

func NewUploader() *Uploader {
	return &Uploader{}
}
