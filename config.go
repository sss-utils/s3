package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type TimeSettings struct {
	LiveTimeDuration time.Duration
}

type Settings struct {
	Url          string
	Region       string
	Bucket       string
	AccessKey    string
	SecretKey    string
	UsePathStyle bool

	Upload   *TimeSettings
	Download *TimeSettings
}

type conf struct {
	client   *s3.Client
	settings Settings
}

func New(settings Settings) Config {
	if settings.Region == "" {
		settings.Region = "ru-1"
	}
	staticProvider := credentials.NewStaticCredentialsProvider(
		settings.AccessKey,
		settings.SecretKey,
		"",
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(settings.Region),
		config.WithBaseEndpoint(settings.Url),
		config.WithCredentialsProvider(staticProvider),
	)

	if err != nil {
		panic(fmt.Errorf("ошибка загрузки конфигурации: %v", err))
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = settings.UsePathStyle
	})

	if settings.Upload == nil {
		settings.Upload = &TimeSettings{
			LiveTimeDuration: 15 * time.Minute,
		}
	}

	if settings.Download == nil {
		settings.Download = &TimeSettings{
			LiveTimeDuration: 15 * time.Minute,
		}
	}

	return conf{
		s3Client,
		settings,
	}
}
