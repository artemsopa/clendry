package storage

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
)

type FileStorage struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewFileStorage(client *minio.Client, bucket, endpoint string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (fs *FileStorage) GetLink(fileName string) string {
	return fs.generateFileURL(fileName)
}

func (fs *FileStorage) Upload(ctx context.Context, input UploadInput) error {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := fs.client.PutObject(ctx, fs.bucket, input.Name, input.File, input.Size, opts)
	return err
}

func (fs *FileStorage) Delete(ctx context.Context, object string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: false,
	}
	err := fs.client.RemoveObject(ctx, fs.bucket, object, opts)
	return err
}

// DigitalOcean Spaces URL format.
func (fs *FileStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("http://%s/%s/%s", fs.endpoint, fs.bucket, filename)
}
