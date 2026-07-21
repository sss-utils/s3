package s3

import (
	"context"
	"fmt"
	"mime"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3Lib "github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c conf) GetPresignedUploadURL(objectKey string, contentType *string) (string, error) {
	if contentType == nil {
		ext := filepath.Ext(objectKey)
		contentType = new(mime.TypeByExtension(ext))
	}

	presignClient := s3Lib.NewPresignClient(c.client)

	input := &s3Lib.PutObjectInput{
		Bucket:      &c.settings.Bucket,
		Key:         &objectKey,
		ContentType: aws.String(*contentType),
	}

	presignedReq, err := presignClient.PresignPutObject(context.TODO(), input, func(o *s3Lib.PresignOptions) {
		o.Expires = c.settings.Upload.LiveTimeDuration

	})

	if err != nil {
		return "", fmt.Errorf("не удалось подписать ссылку: %v", err)
	}

	return presignedReq.URL, nil
}

func (c conf) GetPresignedUploadURLX(objectKey string, contentType *string) string {
	url, err := c.GetPresignedUploadURL(objectKey, contentType)

	if err != nil {
		panic(err)
	}

	return url
}
