package cloudinary

import (
	"net/url"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitCloudinary() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	parsedURL, err := url.Parse(cloudinaryURL)
	if err != nil {
		return nil, err
	}

	apiKey := parsedURL.User.Username()
	apiSecret, _ := parsedURL.User.Password()
	cloudName := parsedURL.Host

	return cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
}
