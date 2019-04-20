package ingest

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// createSession Creates the sessions needed to work with the AWS SDk
func createSession(accountFlag string, regionFlag string) *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(regionFlag)},
		Profile: accountFlag,
	}))
	svc := ec2.New(sess)
	return svc
}

func createCloudwatchSession(accountFlag string, regionFlag string) *cloudwatch.CloudWatch {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(regionFlag)},
		Profile: accountFlag,
	}))
	svc := cloudwatch.New(sess)
	return svc
}
