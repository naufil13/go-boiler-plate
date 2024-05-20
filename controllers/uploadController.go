package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func MultiUpload(c *gin.Context) {

	// Multipart form
	form, err := c.MultipartForm()
	files := form.File["upload"]

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to upload files",
		})
		return
	}

	for _, file := range files {
		log.Println(file.Filename)
		c.SaveUploadedFile(file, os.Getenv("ASSET_PATH")+"file-upload/"+file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "file uploaded",
		"data":    len(files),
	})
}

func MultiUploadS3(c *gin.Context) {
	dirPath := os.Getenv("AWS_S3_BUCKET_PATH") + "/wallet/"
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get multipart form"})
		return
	}

	files := form.File["upload"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files to upload"})
		return
	}

	// Load the AWS configuration

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load AWS config"})
		return
	}

	// Create an S3 service client
	svc := s3.NewFromConfig(cfg)

	// Check if the folder exists in S3
	exists, err := folderExistsInS3(svc, dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check if folder exists in S3"})
		return
	}

	log.Println(exists)

	// Create the folder in S3 if it doesn't exist
	if !exists {
		err = createFolderInS3(svc, dirPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create folder in S3"})
			return
		}
	}

	for _, file := range files {
		// Open the file
		f, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer f.Close()

		//fKey := filepath.Join(dirPath, file.Filename)

		fKey := dirPath + file.Filename
		fKey = strings.ReplaceAll(fKey, "\\", "/")
		log.Println(fKey)

		// Upload the file to S3
		_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
			Key:    aws.String(fKey),
			Body:   f,
		})
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file to S3"})
			return
		}

		log.Println("Successfully uploaded file:", file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Files uploaded successfully",
		"data":    len(files),
	})
}

func createFolderInS3(svc *s3.Client, folderName string) error {
	_, err := svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(folderName),
		Body:   nil,
	})
	return err
}

func folderExistsInS3(svc *s3.Client, folderName string) (bool, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(os.Getenv("AWS_S3_BUCKET")),
		Prefix:  aws.String(folderName),
		MaxKeys: aws.Int32(1),
	}

	result, err := svc.ListObjectsV2(context.TODO(), input)
	if err != nil {
		return false, err
	}

	if len(result.Contents) == 0 {
		return false, nil
	}

	return true, nil
}
