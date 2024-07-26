package upload_handler

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	region          = "sgp1"
	s3Endpoint      = "https://sgp1.digitaloceanspaces.com" // e.g., "nyc3.digitaloceanspaces.com"
	s3Bucket        = "oggspace"
	accessKeyID     = "DO00DZ6UNDJWQZCB4WQF"
	secretAccessKey = "2WztxeEdT+1jn3yOeXX7gVbAcCwdXkQ0/GW+UiEgG64"
)

func UploadImage(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Create buffer to store file content
	fileBytes := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBytes, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading file"})
		return
	}

	fileType := http.DetectContentType(fileBytes.Bytes())
	path := uuid.NewString() + ".jpg" // Specify the path inside the bucket where you want to store the file

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
		Endpoint: aws.String(s3Endpoint),
	})
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(s3Bucket),
		Body:          bytes.NewReader(fileBytes.Bytes()),
		ContentLength: aws.Int64(int64(fileBytes.Len())),
		ContentType:   aws.String(fileType),
		Key:           aws.String(path),
		ACL:           aws.String("public-read"),
	}

	_, err = svc.PutObject(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": "https://oggspace.sgp1.digitaloceanspaces.com/" + path})
}
