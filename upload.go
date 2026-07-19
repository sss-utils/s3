package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c conf) GetPresignedUploadURL(objectKey string) (string, error) {
	presignClient := s3.NewPresignClient(c.client)

	input := &s3.PutObjectInput{
		Bucket: &c.settings.Bucket,
		Key:    &objectKey,
	}

	presignedReq, err := presignClient.PresignPutObject(context.TODO(), input, func(o *s3.PresignOptions) {
		o.Expires = c.settings.Upload.LiveTimeDuration
	})
	if err != nil {
		return "", fmt.Errorf("не удалось подписать ссылку: %v", err)
	}

	return presignedReq.URL, nil
}

func (c conf) GetPresignedUploadURLX(objectKey string) string {
	url, err := c.GetPresignedUploadURL(objectKey)

	if err != nil {
		panic(err)
	}

	return url
}
