package storage

import (
	"context"
	"io"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

type Provider interface {
	GetLink(fileName string) string
	Upload(ctx context.Context, input UploadInput) error
	Delete(ctx context.Context, object string) error
}
