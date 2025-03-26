package cloudinary

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, filename string) (string, error) {
	cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return "", fmt.Errorf("failed to init Cloudinary: %v", err)
	}
	safeFilename := sanitizeFilename(filename)

	uploadParams := uploader.UploadParams{
		PublicID: safeFilename,
	}

	resp, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", fmt.Errorf("upload error: %v", err)
	}

	return resp.SecureURL, nil
}

func sanitizeFilename(filename string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Replace spaces and special characters with dashes
	reg := regexp.MustCompile(`[^\w\-]`)
	safe := reg.ReplaceAllString(name, "-")

	// Remove multiple dashes
	safe = regexp.MustCompile(`-+`).ReplaceAllString(safe, "-")

	return strings.ToLower(safe)
}
