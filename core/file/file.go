package file

import "os"

type (
	FileURL string
	Filer   interface {
		UploadFile(file os.File) (FileURL, error)
		LoadFile(url FileURL) (os.File, error)
	}
)
