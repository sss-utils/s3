package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c conf) GetPresignedDownloadURL(objectKey string) (string, error) {
	presignClient := s3.NewPresignClient(c.client)

	input := &s3.GetObjectInput{
		Bucket:                     &c.settings.Bucket,
		Key:                        &objectKey,
		ResponseContentDisposition: new(fmt.Sprintf("attachment; filename=\"%s\"", objectKey)),
	}

	presignedReq, err := presignClient.PresignGetObject(context.TODO(), input, func(o *s3.PresignOptions) {
		o.Expires = c.settings.Download.LiveTimeDuration
	})
	if err != nil {
		return "", fmt.Errorf("не удалось подписать ссылку на скачивание: %v", err)
	}

	return presignedReq.URL, nil
}

func (c conf) GetPresignedDownloadURLX(objectKey string) string {
	url, err := c.GetPresignedDownloadURL(objectKey)

	if err != nil {
		panic(err)
	}

	return url
}
