package http

import (
	"bytes"

	"github.com/Crodu/CasamentoBackend/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

const (
	AWS_S3_REGION = "us-east-1"
	S3_BUCKET     = "giftphotos"
)

func connectAWS(key string, secret string) *session.Session {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			key,
			secret,
			""),
	})
	if err != nil {
		panic(err)
	}
	return s
}

func UploadFileToS3(c *gin.Context) {
	S3AccessKey := c.MustGet("config").(config.Config).S3AccessKey
	S3SecretKey := c.MustGet("config").(config.Config).S3SecretKey

	// Get the S3 bucket name from the config

	// Create a new S3 client
	session := connectAWS(S3AccessKey, S3SecretKey)

	s3Client := s3.New(session)

	_, err := session.Config.Credentials.Get()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to get AWS credentials, " + err.Error(),
		})
		return
	}

	// Create a buffer to hold the file content
	fileBuffer := &bytes.Buffer{}
	fileName := c.Query("file_name")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get file from request",
		})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer openedFile.Close()

	_, err = fileBuffer.ReadFrom(openedFile)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	// Upload the file to S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileBuffer.Bytes()),
	})

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to upload file to S3, " + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "File uploaded successfully",
		"file_url": "https://" + S3_BUCKET + ".s3.amazonaws.com/" + fileName,
	})
	// c.String(200, "File uploaded successfully to S3 bucket: %s", S3_BUCKET)
	// c.String(200, "File URL: https://%s.s3.amazonaws.com/%s", S3_BUCKET, fileName)
	return
}
