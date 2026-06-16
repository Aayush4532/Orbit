package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"Orbit/configs"
	"Orbit/internal/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	MaxProfileSize = 2 * 1024 * 1024 // 2MB
	MaxEventSize   = 5 * 1024 * 1024 // 5MB
	MaxProductSize = 4 * 1024 * 1024 // 4MB
)

type UploadPhotoToR2 interface {
	UploadPhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string) (string, error)
	DeletePhoto(ctx context.Context, folder string, filename string) error
	UpdatePhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string, oldFolder string, oldFilename string) (string, error)
}

type Profile struct{}
type Event struct{}
type Product struct{}

func validateAndDecode(fileHeader *multipart.FileHeader, maxSize int64) (image.Image, string, error) {
	if fileHeader == nil {
		return nil, "", errors.New("multipart file header is missing")
	}

	if fileHeader.Size > maxSize {
		return nil, "", fmt.Errorf("file too large: maximum allowed size is %d MB", maxSize/(1024*1024))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, "", fmt.Errorf("failed to read file header: %w", err)
	}

	contentType := http.DetectContentType(buffer[:n])
	if contentType != "image/jpeg" && contentType != "image/png" {
		return nil, "", errors.New("invalid file standard: only JPG, JPEG, and PNG formats are allowed")
	}

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, "", fmt.Errorf("failed to reset file pointer: %w", err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, "", errors.New("corrupted image content or encoding failure")
	}

	return img, contentType, nil
}

func uploadToR2Internal(ctx context.Context, folder, filename, contentType string, body io.Reader) (string, error) {
	r2Cfg := configs.LoadConfig().R2 // Note: Optimization tip - ensure this configuration look-up is cached in configs package memory
	client := db.GetR2Client()

	key := fmt.Sprintf("%s/%s", folder, filename)

	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2Cfg.Bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("cloudflare r2 driver error: %w", err)
	}

	return fmt.Sprintf("%s/%s", r2Cfg.Domain, key), nil
}

func deleteFromR2Internal(ctx context.Context, folder, filename string) error {
	r2Cfg := configs.LoadConfig().R2
	key := fmt.Sprintf("%s/%s", folder, filename)

	_, err := db.GetR2Client().DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(r2Cfg.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object from r2: %w", err)
	}
	return nil
}

func (p *Profile) UploadPhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string) (string, error) {
	img, contentType, err := validateAndDecode(fileHeader, MaxProfileSize)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	var filename string

	if contentType == "image/png" {
		filename = fmt.Sprintf("%s-%d.png", id, time.Now().UnixNano())
		err = png.Encode(buf, img)
	} else {
		filename = fmt.Sprintf("%s-%d.jpg", id, time.Now().UnixNano())
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 82})
		contentType = "image/jpeg"
	}

	if err != nil {
		return "", fmt.Errorf("profile encoding pipeline failure: %w", err)
	}

	return uploadToR2Internal(ctx, "profile", filename, contentType, buf)
}

func (p *Profile) DeletePhoto(ctx context.Context, folder string, filename string) error {
	return deleteFromR2Internal(ctx, folder, filename)
}

func (p *Profile) UpdatePhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string, oldFolder string, oldFilename string) (string, error) {
	newURL, err := p.UploadPhoto(ctx, fileHeader, id)
	if err != nil {
		return "", err
	}

	if oldFilename != "" && oldFolder != "" {
		if delErr := p.DeletePhoto(ctx, oldFolder, oldFilename); delErr != nil {
			log.Printf("WARN: Orphan file alert - failed to delete old profile photo %s/%s: %v", oldFolder, oldFilename, delErr)
		}
	}

	return newURL, nil
}

func (e *Event) UploadPhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string) (string, error) {
	img, contentType, err := validateAndDecode(fileHeader, MaxEventSize)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	var filename string

	if contentType == "image/png" {
		filename = fmt.Sprintf("%s-%d.png", id, time.Now().UnixNano())
		err = png.Encode(buf, img)
	} else {
		filename = fmt.Sprintf("%s-%d.jpg", id, time.Now().UnixNano())
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 85})
		contentType = "image/jpeg"
	}

	if err != nil {
		return "", fmt.Errorf("event encoding pipeline failure: %w", err)
	}

	return uploadToR2Internal(ctx, "event-banner", filename, contentType, buf)
}

func (e *Event) DeletePhoto(ctx context.Context, folder string, filename string) error {
	return deleteFromR2Internal(ctx, folder, filename)
}

func (e *Event) UpdatePhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string, oldFolder string, oldFilename string) (string, error) {
	newURL, err := e.UploadPhoto(ctx, fileHeader, id)
	if err != nil {
		return "", err
	}

	if oldFilename != "" && oldFolder != "" {
		if delErr := e.DeletePhoto(ctx, oldFolder, oldFilename); delErr != nil {
			log.Printf("WARN: Orphan file alert - failed to delete old event banner %s/%s: %v", oldFolder, oldFilename, delErr)
		}
	}

	return newURL, nil
}

func (pr *Product) UploadPhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string) (string, error) {
	img, contentType, err := validateAndDecode(fileHeader, MaxProductSize)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	var filename string

	if contentType == "image/png" {
		filename = fmt.Sprintf("%s-%d.png", id, time.Now().UnixNano())
		err = png.Encode(buf, img)
	} else {
		filename = fmt.Sprintf("%s-%d.jpg", id, time.Now().UnixNano())
		err = jpeg.Encode(buf, img, &jpeg.Options{Quality: 80})
		contentType = "image/jpeg"
	}

	if err != nil {
		return "", fmt.Errorf("product pipeline compression error: %w", err)
	}

	return uploadToR2Internal(ctx, "product", filename, contentType, buf)
}

func (pr *Product) DeletePhoto(ctx context.Context, folder string, filename string) error {
	return deleteFromR2Internal(ctx, folder, filename)
}

func (pr *Product) UpdatePhoto(ctx context.Context, fileHeader *multipart.FileHeader, id string, oldFolder string, oldFilename string) (string, error) {
	newURL, err := pr.UploadPhoto(ctx, fileHeader, id)
	if err != nil {
		return "", err
	}

	if oldFilename != "" && oldFolder != "" {
		if delErr := pr.DeletePhoto(ctx, oldFolder, oldFilename); delErr != nil {
			log.Printf("WARN: Orphan file alert - failed to delete old product image %s/%s: %v", oldFolder, oldFilename, delErr)
		}
	}

	return newURL, nil
}
