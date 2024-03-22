package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type ImageSvcInter interface {
	PostImage(c *gin.Context)
}

type ImageController struct {
}

func NewImageController() ImageSvcInter {
	return &ImageController{}
}

func (ic *ImageController) PostImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(file.Filename)
	fmt.Printf("Type of f: %T\n", file.Filename)

	url, err := uploadFile(file)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"imageUrl": url,
	})
}

func uploadFile(file *multipart.FileHeader) (string, error) {
	// Retrieve configuration values
	accessKey := viper.GetString("S3_ID")
	secretKey := viper.GetString("S3_SECRET_KEY")
	baseUrl := viper.GetString("S3_BASE_URL")
	region := "ap-southeast-1"

	// Create a new AWS session with the retrieved credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	// Create an S3 client
	svc := s3.New(sess)
	fileName := uuid.New().String() + filepath.Ext(file.Filename) // Generate UUID as file name

	// Prepare the parameters for the PutObject request
	params := &s3.PutObjectInput{
		Bucket: aws.String(baseUrl),
		Key:    aws.String(fileName),
		ACL:    aws.String("public-read"),
		Body:   src,
	}

	// Upload file to S3
	_, err = svc.PutObject(params)
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", baseUrl, fileName)

	return url, nil
}