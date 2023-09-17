package uploader

import (
	"context"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

type IUploadHelper interface {
	UploadFile(ctx context.Context, bucketname, filename, contentType string, fileContent io.ReadSeeker) error
}

type uploadHelper struct {
	s3 *s3.S3
}

func NewUploadHelper(awsRegion string) (*uploadHelper, error) {
	region := viper.GetString("region")
	creds := credentials.NewStaticCredentials(viper.GetString("accessKey"), viper.GetString("secretKey"), viper.GetString("token"))
	_, err := creds.Get()
	if err != nil {
		log.Fatal(err)
	}
	// Buat sesi AWS
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	// Inisialisasi Layanan AWS
	svc := s3.New(sess)

	return &uploadHelper{
		s3: svc,
	}, nil
}

func (s3Helper *uploadHelper) UploadFile(ctx context.Context, bucketname, filename, contentType string, fileContent io.ReadSeeker) error {
	// Lakukan upload ke S3
	_, err := s3Helper.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketname),
		Key:         aws.String(filename),
		ContentType: aws.String(contentType),
		Body:        fileContent,
	})
	return err
}
