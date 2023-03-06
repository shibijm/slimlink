package ports

import (
	"net/http"
)

type HttpFileSystem interface {
	GetFile(path string) (http.File, error)
}
