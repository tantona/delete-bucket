package bucket

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	log "github.com/sirupsen/logrus"
)

type Bucket struct {
	Name     string
	Region   string
	S3Client s3iface.S3API
}

func (b *Bucket) deleteObjects(t string, objects []*s3.ObjectIdentifier) error {
	log.Infof("attempting to delete %d %s", len(objects), t)
	if _, err := b.S3Client.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String(b.Name),
		Delete: &s3.Delete{Objects: objects, Quiet: aws.Bool(false)},
	}); err != nil {
		return fmt.Errorf("unable to delete objects in bucket %s: %v", b.Name, err)
	}

	log.Debugf("removed %d %s", len(objects), t)
	return nil
}

func (b *Bucket) Delete() error {
	listObjectVersionsResponse, err := b.S3Client.ListObjectVersions(&s3.ListObjectVersionsInput{
		Bucket: aws.String(b.Name),
	})

	if err != nil {
		return fmt.Errorf("unable to list object versions in bucket %s: %v", b.Name, err)
	}

	if len(listObjectVersionsResponse.Versions) > 0 {
		o := []*s3.ObjectIdentifier{}
		for _, c := range listObjectVersionsResponse.Versions {
			o = append(o, &s3.ObjectIdentifier{Key: c.Key, VersionId: c.VersionId})
		}

		if err := b.deleteObjects("Versions", o); err != nil {
			return fmt.Errorf("unable to delete Object Versions bucket %s: %v", b.Name, err)
		}
	}

	if len(listObjectVersionsResponse.DeleteMarkers) > 0 {
		o := []*s3.ObjectIdentifier{}
		for _, c := range listObjectVersionsResponse.DeleteMarkers {
			o = append(o, &s3.ObjectIdentifier{Key: c.Key, VersionId: c.VersionId})
		}

		if err := b.deleteObjects("DeleteMarkers", o); err != nil {
			return fmt.Errorf("unable to delete Delete Markers in bucket %s: %v", b.Name, err)
		}
	}

	if _, err := b.S3Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(b.Name)}); err != nil {
		return fmt.Errorf("unable to remove bucket %s: %v", b.Name, err)
	}

	log.Infof("Deleted bucket: %s", b.Name)

	return nil
}

func New(client s3iface.S3API, name, region string) *Bucket {
	return &Bucket{
		Name:     name,
		Region:   region,
		S3Client: client,
	}
}
