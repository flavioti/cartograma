package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func MyListObjects(client s3.Client, bucketName string) {
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(os.Getenv("AWS_PROFILE")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output1, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output1.Buckets {
		log.Printf("name=%s", aws.ToString(object.Name))

		MyListObjects(*client, *object.Name)

	}

}
