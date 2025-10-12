package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/storage"
)
var gcsClient *storage.Client;

func InitiGCS() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	gcsClient = client
	fmt.Println("GCS client initialized")
}

func GenerateV4PutObjectSignedURL( bucket, object, method string) (string, error) {
	// bucket := "bucket-name"
	// object := "object-name"

	
	defer gcsClient.Close()
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: method,
		GoogleAccessID:os.Getenv("GCP_CLIENT_EMAIL"),
		PrivateKey: []byte(os.Getenv("GCP_PRIVATE_KEY")),
	}

	if method == "PUT" {
		opts.Headers = []string{"Content-Type:application/octet-stream"}
		opts.Expires = time.Now().Add(15 * time.Minute)
	} else {
		opts.Expires = time.Now().Add(2 * time.Hour)
	}

	u, err := gcsClient.Bucket(bucket).SignedURL(object, opts)
	if err != nil {
		return "", fmt.Errorf("Bucket(%q).SignedURL: %w", bucket, err)
	}
	return u, nil
}



// we dont need service account initialization for  genreateing the signed url\

// we need it for delete and other operations from backend
