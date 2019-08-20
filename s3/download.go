package s3

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// DownloadFiles download files recursively from s3 (based on bucket + base folder + pattern + download location)
func DownloadFiles(profile string, bucket string, baseFolder string, pattern string, downloadLocation string) error {
	fmt.Println("Getting files from the s3 bucket :", bucket)
	fmt.Println("And will download them to :", downloadLocation)
	sess, err := MakeSession(profile)
	if err != nil {
		return err
	}
	numberOfObjects, objErr := DownloadBucketObjects(sess, bucket, baseFolder, pattern, downloadLocation)
	if objErr != nil {
		return objErr
	}
	fmt.Println(fmt.Sprintf("Number of downloaded objects: %d", numberOfObjects))
	return nil
}

// MakeSession create an aws client session, based on default or custom settings
func MakeSession(profile string) (*session.Session, error) {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	var sess *session.Session
	var err error
	if true {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region:      aws.String("eu-west-1"),
				Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
			},
			Profile: profile,
		})
	} else {
		sess, err = session.NewSessionWithOptions(session.Options{
			Config: aws.Config{
				Region: aws.String("eu-west-1"),
			},
			Profile: profile,
		})
	}
	if err != nil {
		return nil, err
	}

	return sess, nil
}

// DownloadBucketObjects download objects recursively from s3 (based on bucket + base folder + pattern + download location)
func DownloadBucketObjects(sess *session.Session, bucket string, baseFolder string, pattern string, downloadLocation string) (int, error) {
	query := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(baseFolder),
	}
	svc := s3.New(sess)
	numberOfObjects := 0
	// Flag used to check if we need to go further
	truncatedListing := true

	for truncatedListing {
		resp, err := svc.ListObjectsV2(query)
		if err != nil {
			return 0, err
		}
		for _, key := range resp.Contents {
			if len(pattern) > 0 {
				if strings.Contains(*key.Key, pattern) {
					objErr := DownloadObject(key, svc, bucket, downloadLocation)
					if objErr != nil {
						return numberOfObjects, objErr
					}
					numberOfObjects++
				}
			} else {
				objErr := DownloadObject(key, svc, bucket, downloadLocation)
				if objErr != nil {
					return numberOfObjects, objErr
				}
				numberOfObjects++
			}
		}
		query.ContinuationToken = resp.NextContinuationToken
		truncatedListing = *resp.IsTruncated
	}
	return numberOfObjects, nil
}

// DownloadObject download a specific s3 object to a predefined location
func DownloadObject(key *s3.Object, s3Client *s3.S3, bucket string, downloadLocation string) error {
	destFilename := *key.Key
	if strings.HasSuffix(*key.Key, "/") {
		fmt.Println("Got a directory")
		return nil
	}
	fmt.Println(*key.Key)
	if strings.Contains(*key.Key, "/") {
		var dirTree string
		// split
		s3FileFullPathList := strings.Split(*key.Key, "/")
		fmt.Println(s3FileFullPathList)
		fmt.Println("destFilename " + destFilename)
		for _, dir := range s3FileFullPathList[:len(s3FileFullPathList)-1] {
			dirTree = path.Join(dirTree, dir)
		}
		os.MkdirAll(path.Join(downloadLocation, dirTree), 0775)
	}
	out, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    key.Key,
	})
	if err != nil {
		return err
	}
	defer out.Body.Close()
	destFilePath := path.Join(downloadLocation, destFilename)
	destFile, err := os.Create(destFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()
	bytes, err := io.Copy(destFile, out.Body)
	if err != nil {
		return err
	}
	fmt.Printf("File %s contanin %d bytes\n", destFilePath, bytes)
	return nil
}
