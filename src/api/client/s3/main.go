package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/rodrigoherera/know-vegan-service/src/api/config"
)

var (
	GetSession *s3.S3
)

func init() {
	InitS3()
}

func InitS3() {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSREGION),
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	GetSession = s3.New(sess)
}
