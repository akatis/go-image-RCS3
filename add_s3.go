package go_image_RCS3

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func AddS3(imageFile, imagePath string) error {
	//START AWS SESSION
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials(S3_ACCESS_KEY, S3_SECRET_KEY, ""),
	})
	if err != nil {
		return err
	}

	//S3 CLIENT
	s3Client := s3.New(sess)

	//filePath := "./test/denemePng.jpg"

	imageBase64, err := base64.StdEncoding.DecodeString(imageFile)
	if err != nil {
		return err
	}

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(S3_BUCKET),
		Key:         aws.String(S3_OBJECT_KEY + imagePath),
		Body:        bytes.NewReader(imageBase64),
		ContentType: aws.String("image/jpg"),
	})
	if err != nil {
		return err
	}

	return nil
}
