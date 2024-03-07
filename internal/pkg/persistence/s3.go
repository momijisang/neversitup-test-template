package persistence

import (
	"io"
	"net/http"
	"neversitup-test-template/internal/pkg/config"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Repository struct{}

var s3Repository *S3Repository

func S3R() *S3Repository {
	if s3Repository == nil {
		s3Repository = &S3Repository{}
	}
	return s3Repository
}

func (r *S3Repository) FileToS3(filename, mimeType string) (string, error) {
	s3Url, err := r.Upload("contentimage", filename, mimeType)
	if err != nil {
		return "", err
	}
	_ = os.Remove("./tmp/" + filename)

	return s3Url, nil
}

func (r *S3Repository) Download(url, filename string) error {
	// Send a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the output file
	file, err := os.Create("./tmp/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body to the output file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (r *S3Repository) Upload(MediaID string, Filename string, ContentType string) (string, error) {
	f, err := os.Open("./tmp/" + Filename)
	if err != nil {
		return "", err
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.Config.S3Server.Region),
		Credentials: credentials.NewStaticCredentials(config.Config.S3Server.Key, config.Config.S3Server.Secret, "")}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(config.Config.S3Server.Bucket),
		Key:         aws.String(MediaID + "/" + Filename),
		Body:        f,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(ContentType),
	})

	if err != nil {
		return "", err
	}
	return result.Location, nil
}

func (r *S3Repository) DeleteFile(Filename string) error {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(config.Config.S3Server.Region),
		Credentials: credentials.NewStaticCredentials(config.Config.S3Server.Key, config.Config.S3Server.Secret, "")}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.Config.S3Server.Bucket),
		Key:    aws.String(Filename),
	})

	return err
}
