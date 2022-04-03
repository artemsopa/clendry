package files

import (
	"context"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
)

type Files interface {
	GetObjectLink(fileName string) string
	UploadObject(ctx context.Context, folder string, file service.File) (string, error)
	RemoveObject(ctx context.Context, object string) error
	RemoveFile(filename string)
}
