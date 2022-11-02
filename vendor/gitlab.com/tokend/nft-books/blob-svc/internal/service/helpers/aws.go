package helpers

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
)
import "github.com/aws/aws-sdk-go/aws/session"

func NewAWSSession(config *config.AWSConfig) *session.Session {
	awsSessionConfig := aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID,
			config.SecretKeyID,
			""),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(config.SslDisable),
		S3ForcePathStyle: aws.Bool(config.ForcePathStyle),
	}

	if config.Endpoint != "" {
		awsSessionConfig.Endpoint = aws.String(config.Endpoint)
	}

	return session.Must(session.NewSession(&awsSessionConfig))
}

func UploadFile(file multipart.File, key string, config *config.AWSConfig) error {
	awsSession := NewAWSSession(config)
	uploader := s3manager.NewUploader(awsSession)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func GetUrl(key string, config *config.AWSConfig) (string, error) {
	awsSession := NewAWSSession(config)
	service := s3.New(awsSession)

	req, _ := service.GetObjectRequest(&s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(config.Bucket),
	})

	return req.Presign(config.Expiration)
}

func DeleteFile(key string, config *config.AWSConfig) error {
	awsSession := NewAWSSession(config)
	service := s3.New(awsSession)

	_, err := service.DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(config.Bucket),
	})
	if err != nil {
		return err
	}

	err = service.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(config.Bucket),
	})

	return err
}

func IsKeyExists(key string, config *config.AWSConfig) (bool, error) {
	awsSession := NewAWSSession(config)
	service := s3.New(awsSession)

	_, err := service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound":
				return false, nil
			default:
				return false, err
			}
		}

		return false, err
	}

	return true, nil
}
