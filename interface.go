package s3

type Config interface {
	GetPresignedDownloadURL(objectKey string) (string, error)
	GetPresignedDownloadURLX(objectKey string) string

	GetPresignedUploadURL(objectKey string) (string, error)
	GetPresignedUploadURLX(objectKey string) string
}
