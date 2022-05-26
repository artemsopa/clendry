package files

import (
	"context"

	"github.com/artomsopun/clendry/clendry-api/internal/domain"
)

type File struct {
	Title       string
	Size        int64
	ContentType string
	Type        domain.FileType
}

type Files interface {
	GetObjectLink(fileName string) string
	UploadObject(ctx context.Context, userID string, folder string, file File) (string, string, error)
	RemoveObject(ctx context.Context, object string) error
	RemoveFile(filename string)
}
