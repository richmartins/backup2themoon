package main

import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func main() {
	// Load the Shared AWS Configuration from access key and secret key
	cfg, err := config.LoadDefaultConfig(context.TODO(), 
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("ACCESS_KEY", "SECRET_ACCESS_KEY", "")),
		config.WithRegion("us-east-1"),
		// give s3.pub1.infomaniak.cloud as endpoint, if you want to use Infomaniak S3
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "https://s3.pub1.infomaniak.cloud",
				SigningRegion: "us-east-1",
			}, nil
		}),
	))

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	//list buckets in the account
	buckets, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("buckets:")
	for _, bucket := range buckets.Buckets {
		log.Println(aws.ToString(bucket.Name))
	}


	// Set up parameters for listing objects
    input := &s3.ListObjectsV2Input{
        Bucket: aws.String("backup2themoon-test"),
    }

    // List all objects in the bucket
	result, err := client.ListObjectsV2(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("objects:")
	for _, object := range result.Contents {
		log.Println(aws.ToString(object.Key))
	}

}