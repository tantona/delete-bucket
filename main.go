package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
	"github.com/tantona/delete-bucket/bucket"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		verbose    = kingpin.Flag("verbose", "enable verbose logging").Default("false").Short('v').Bool()
		profile    = kingpin.Flag("profile", "aws profile name").Short('p').Default("").String()
		region     = kingpin.Flag("region", "aws region").Short('r').Default("us-east-1").String()
		bucketName = kingpin.Arg("bucket name", "name of s3 bucket").Required().String()
	)
	kingpin.Parse()

	log.SetLevel(log.InfoLevel)
	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: region},
		Profile: *profile,
	})
	if err != nil {
		log.Fatal(err)
	}

	bucketRegion, err := s3manager.GetBucketRegionWithClient(context.Background(), s3.New(sess), *bucketName)
	if err != nil {
		log.Fatal(err)
	}

	sess, err = session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(bucketRegion)},
		Profile: *profile,
	})
	if err != nil {
		log.Fatal(err)
	}

	b := bucket.New(s3.New(sess), *bucketName, bucketRegion)

	if err := b.Delete(); err != nil {
		log.Fatal(err)
	}

	log.Infof("Successfully Deleted Bucket s3://%s (%s)", *bucketName, bucketRegion)
}
