package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	endpoint := "localhost:9000"
	accessKeyID := "user"
	secretAccessKey := "password"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	opts := minio.ListObjectsOptions{
		Recursive: true,
		Prefix:    "",
	}

	for object := range minioClient.ListObjects(context.Background(), "first", opts) {
		fmt.Println(object.Key, object.Size, object.Metadata)
		minioClient.FGetObject(context.Background(),
			"first", object.Key,
			"tmp/"+object.Key,
			minio.GetObjectOptions{})
	}

}
