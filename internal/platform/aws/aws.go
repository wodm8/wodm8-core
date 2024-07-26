package aws

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime"
	"path/filepath"
)

var region = "us-east-1"

type Awsservice interface {
	Load() (aws.Config, error)
	S3(cfg aws.Config) (*s3.Client, error)
	SaveObjectInBucket(clt *s3.Client) error
}

type AwsserviceImpl struct {
	ctx context.Context
}

func NewAwsServiceImpl(ctx context.Context) *AwsserviceImpl {
	return &AwsserviceImpl{ctx: ctx}
}

func (s *AwsserviceImpl) Load() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(s.ctx, config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
		return aws.Config{}, err
	}
	return cfg, err
}

func (s *AwsserviceImpl) S3(cfg aws.Config) *s3.Client {
	client := s3.NewFromConfig(cfg)
	return client
}

func (s *AwsserviceImpl) SaveObjectInBucket(clt *s3.Client, bucket, key string, file []byte) (string, error) {

	contentType := mime.TypeByExtension(filepath.Ext(key))
	fmt.Printf("this is the type of file, %s\n", contentType)
	if contentType == "" {
		contentType = "application/octet-stream" // valor por defecto si no se puede determinar
	}
	fmt.Printf("Bucket: %s, Key: %s\n", bucket, key)
	_, err := clt.PutObject(s.ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(file),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		fmt.Printf("error loading image to s3, %v", err)
		return "", err
	}
	objectURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)

	return objectURL, nil
}
