package errors

import "errors"

var (
	ResourceNotFound      = errors.New("resource was not found")
	CredentialsDoNotMatch = errors.New("credentials do not match")

	FileIsTooBig = errors.New("file size is too big")
)
