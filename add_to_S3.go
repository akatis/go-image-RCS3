package go_image_RCS3

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

type S3Config struct {
	S3_ACCESS_KEY string
	S3_SECRET_KEY string
	S3_REGION     string
	S3_BUCKET     string
	S3_OBJECT_KEY string
}

func New(config *S3Config) *S3Config {
	return config
}

func (s3Config *S3Config) AddS3(imageFile, imagePath string) error {
	//START AWS SESSION
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.S3_REGION),
		Credentials: credentials.NewStaticCredentials(s3Config.S3_ACCESS_KEY, s3Config.S3_SECRET_KEY, ""),
	})
	if err != nil {
		return err
	}

	//S3 CLIENT
	s3Client := s3.New(sess)

	//base64 format controller
	index := strings.Index(imageFile, ";base64,")
	if index < 0 {
		imageBase64, err := base64.StdEncoding.DecodeString(imageFile)
		if err != nil {
			return err
		}

		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(s3Config.S3_BUCKET),
			Key:         aws.String(s3Config.S3_OBJECT_KEY + imagePath),
			Body:        bytes.NewReader(imageBase64),
			ContentType: aws.String("image/jpg"),
		})
		if err != nil {
			return err
		}
	}
	imgExt := imageFile[11:index]
	imageBase64, err := base64.StdEncoding.DecodeString(imageFile[index+8:])
	if err != nil {
		return err
	}

	switch imgExt {
	case "png":
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(s3Config.S3_BUCKET),
			Key:         aws.String(s3Config.S3_OBJECT_KEY + imagePath),
			Body:        bytes.NewReader(imageBase64),
			ContentType: aws.String("image/png"),
		})
		if err != nil {
			return err
		}
	case "jpeg", "jpg":
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(s3Config.S3_BUCKET),
			Key:         aws.String(s3Config.S3_OBJECT_KEY + imagePath),
			Body:        bytes.NewReader(imageBase64),
			ContentType: aws.String("image/jpg"),
		})
		if err != nil {
			return err
		}
	case "svg+xml":
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(s3Config.S3_BUCKET),
			Key:         aws.String(s3Config.S3_OBJECT_KEY + imagePath),
			Body:        bytes.NewReader(imageBase64),
			ContentType: aws.String("image/svg+xml"),
		})
		if err != nil {
			return err
		}
	case "webp":
		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(s3Config.S3_BUCKET),
			Key:         aws.String(s3Config.S3_OBJECT_KEY + imagePath),
			Body:        bytes.NewReader(imageBase64),
			ContentType: aws.String("image/webp"),
		})
		if err != nil {
			return err
		}

	}

	return nil
}

// -------DELETE FROM S3--------
func (s3Config *S3Config) DeleteS3(imagePath string) error {
	//START AWS SESSION
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Config.S3_REGION),
		Credentials: credentials.NewStaticCredentials(s3Config.S3_ACCESS_KEY, s3Config.S3_SECRET_KEY, ""),
	})
	if err != nil {
		return err
	}

	//S3 CLIENT
	s3Client := s3.New(sess)

	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s3Config.S3_BUCKET),
		Key:    aws.String(s3Config.S3_OBJECT_KEY + imagePath),
	})
	if err != nil {
		return err
	}

	return nil
}
